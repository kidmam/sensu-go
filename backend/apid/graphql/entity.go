package graphql

import (
	"sort"
	"time"

	"github.com/sensu/sensu-go/backend/apid/graphql/globalid"
	"github.com/sensu/sensu-go/backend/apid/graphql/schema"
	"github.com/sensu/sensu-go/graphql"
	"github.com/sensu/sensu-go/types"
	"github.com/sensu/sensu-go/util/strings"
)

var _ schema.EntityFieldResolvers = (*entityImpl)(nil)
var _ schema.SystemFieldResolvers = (*systemImpl)(nil)
var _ schema.NetworkFieldResolvers = (*networkImpl)(nil)
var _ schema.NetworkInterfaceFieldResolvers = (*networkInterfaceImpl)(nil)
var _ schema.DeregistrationFieldResolvers = (*deregistrationImpl)(nil)

//
// Implement EntityFieldResolvers
//

type entityImpl struct {
	schema.EntityAliases
	factory ClientFactory
}

// ID implements response to request for 'id' field.
func (*entityImpl) ID(p graphql.ResolveParams) (string, error) {
	return globalid.EntityTranslator.EncodeToString(p.Source), nil
}

// ExtendedAttributes implements response to request for 'extendedAttributes' field.
func (*entityImpl) ExtendedAttributes(p graphql.ResolveParams) (interface{}, error) {
	entity := p.Source.(*types.Entity)
	return wrapExtendedAttributes(entity.ExtendedAttributes), nil
}

// LastSeen implements response to request for 'executed' field.
func (r *entityImpl) LastSeen(p graphql.ResolveParams) (*time.Time, error) {
	e := p.Source.(*types.Entity)
	return convertTs(e.LastSeen), nil
}

// Events implements response to request for 'events' field.
func (r *entityImpl) Events(p schema.EntityEventsFieldResolverParams) (interface{}, error) {
	src := p.Source.(*types.Entity)
	client := r.factory.NewWithContext(p.Context)

	// fetch
	evs, err := fetchEvents(client, src.Namespace, func(obj *types.Event) bool {
		return obj.Entity.Name == src.Name
	})
	if err != nil {
		return []interface{}{}, err
	}

	// sort records
	sortEvents(evs, p.Args.OrderBy)

	return evs, nil
}

// Related implements response to request for 'related' field.
func (r *entityImpl) Related(p schema.EntityRelatedFieldResolverParams) (interface{}, error) {
	entity := p.Source.(*types.Entity)
	client := r.factory.NewWithContext(p.Context)

	// fetch & omit source
	entities, err := fetchEntities(client, entity.Namespace, func(obj *types.Entity) bool {
		return obj.Name != entity.Name
	})
	if err != nil {
		return []interface{}{}, err
	}

	// sort
	scores := map[int]int{}
	for i, en := range entities {
		matched := strings.Intersect(
			append(en.Subscriptions, en.EntityClass, en.System.Platform),
			append(entity.Subscriptions, entity.EntityClass, entity.System.Platform),
		)
		scores[i] = len(matched)
	}
	sort.Slice(entities, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	// limit
	limit := clampInt(p.Args.Limit, 0, len(entities))
	return entities[0:limit], nil
}

// Status implements response to request for 'status' field.
func (r *entityImpl) Status(p graphql.ResolveParams) (interface{}, error) {
	src := p.Source.(*types.Entity)
	client := r.factory.NewWithContext(p.Context)

	// fetch
	evs, err := fetchEvents(client, src.Namespace, func(obj *types.Event) bool {
		return obj.Entity.Name == src.Name
	})
	if err != nil {
		return 0, err
	}

	// find MAX value
	var st uint32
	for _, ev := range evs {
		if ev.Check == nil {
			continue
		}
		st = maxUint32(ev.Check.Status, st)
	}
	return st, nil
}

// IsSilenced implements response to request for 'isSilenced' field.
func (r *entityImpl) IsSilenced(p graphql.ResolveParams) (bool, error) {
	src := p.Source.(*types.Entity)
	client := r.factory.NewWithContext(p.Context)
	sls, err := fetchSilenceds(client, src.Namespace, filterSilenceByEntity(src))
	return len(sls) > 0, err
}

// Silences implements response to request for 'silences' field.
func (r *entityImpl) Silences(p graphql.ResolveParams) (interface{}, error) {
	src := p.Source.(*types.Entity)
	client := r.factory.NewWithContext(p.Context)
	sls, err := fetchSilenceds(client, src.Namespace, filterSilenceByEntity(src))
	return sls, err
}

func filterSilenceByEntity(src *types.Entity) silencePredicate {
	now := time.Now().Unix()
	return func(obj *types.Silenced) bool {
		if !(obj.Check == "" || obj.Check == "*") || !obj.StartSilence(now) {
			return false
		}
		if strings.InArray(obj.Subscription, src.Subscriptions) {
			return true
		}
		return false
	}
}

// IsTypeOf is used to determine if a given value is associated with the type
func (*entityImpl) IsTypeOf(s interface{}, p graphql.IsTypeOfParams) bool {
	_, ok := s.(*types.Entity)
	return ok
}

//
// Implement SystemFieldResolvers
//

type systemImpl struct {
	schema.SystemAliases
}

// Os implements response to request for 'os' field.
func (r *systemImpl) Os(p graphql.ResolveParams) (string, error) {
	sys := p.Source.(types.System)
	return sys.OS, nil
}

//
// Implement NetworkFieldResolvers
//

type networkImpl struct {
	schema.NetworkAliases
}

//
// Implement NetworkInterfaceFieldResolvers
//

type networkInterfaceImpl struct {
	schema.NetworkInterfaceAliases
}

// Mac implements response to request for 'mac' field.
func (r *networkInterfaceImpl) Mac(p graphql.ResolveParams) (string, error) {
	i := p.Source.(types.NetworkInterface)
	return i.MAC, nil
}

//
// Implement DeregistrationFieldResolvers
//

type deregistrationImpl struct {
	schema.DeregistrationAliases
}

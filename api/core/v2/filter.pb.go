// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: filter.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// EventFilter is a filter specification.
type EventFilter struct {
	// Metadata contains the name, namespace, labels and annotations of the filter
	ObjectMeta `protobuf:"bytes,1,opt,name=metadata,embedded=metadata" json:"metadata,omitempty"`
	// Action specifies to allow/deny events to continue through the pipeline
	Action string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	// Expressions is an array of boolean expressions that are &&'d together
	// to determine if the event matches this filter.
	Expressions []string `protobuf:"bytes,3,rep,name=expressions" json:"expressions"`
	// When indicates a TimeWindowWhen that a filter uses to filter by days & times
	When *TimeWindowWhen `protobuf:"bytes,6,opt,name=when" json:"when,omitempty"`
	// Runtime assets are Sensu assets that contain javascript libraries. They
	// are evaluated within the execution context.
	RuntimeAssets        []string `protobuf:"bytes,8,rep,name=runtime_assets,json=runtimeAssets" json:"runtime_assets"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventFilter) Reset()         { *m = EventFilter{} }
func (m *EventFilter) String() string { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()    {}
func (*EventFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_fa80d2fd62e20657, []int{0}
}
func (m *EventFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EventFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFilter.Merge(dst, src)
}
func (m *EventFilter) XXX_Size() int {
	return m.Size()
}
func (m *EventFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFilter.DiscardUnknown(m)
}

var xxx_messageInfo_EventFilter proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventFilter)(nil), "sensu.core.v2.EventFilter")
}
func (this *EventFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EventFilter)
	if !ok {
		that2, ok := that.(EventFilter)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ObjectMeta.Equal(&that1.ObjectMeta) {
		return false
	}
	if this.Action != that1.Action {
		return false
	}
	if len(this.Expressions) != len(that1.Expressions) {
		return false
	}
	for i := range this.Expressions {
		if this.Expressions[i] != that1.Expressions[i] {
			return false
		}
	}
	if !this.When.Equal(that1.When) {
		return false
	}
	if len(this.RuntimeAssets) != len(that1.RuntimeAssets) {
		return false
	}
	for i := range this.RuntimeAssets {
		if this.RuntimeAssets[i] != that1.RuntimeAssets[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

type EventFilterFace interface {
	Proto() github_com_golang_protobuf_proto.Message
	GetObjectMeta() ObjectMeta
	GetAction() string
	GetExpressions() []string
	GetWhen() *TimeWindowWhen
	GetRuntimeAssets() []string
}

func (this *EventFilter) Proto() github_com_golang_protobuf_proto.Message {
	return this
}

func (this *EventFilter) TestProto() github_com_golang_protobuf_proto.Message {
	return NewEventFilterFromFace(this)
}

func (this *EventFilter) GetObjectMeta() ObjectMeta {
	return this.ObjectMeta
}

func (this *EventFilter) GetAction() string {
	return this.Action
}

func (this *EventFilter) GetExpressions() []string {
	return this.Expressions
}

func (this *EventFilter) GetWhen() *TimeWindowWhen {
	return this.When
}

func (this *EventFilter) GetRuntimeAssets() []string {
	return this.RuntimeAssets
}

func NewEventFilterFromFace(that EventFilterFace) *EventFilter {
	this := &EventFilter{}
	this.ObjectMeta = that.GetObjectMeta()
	this.Action = that.GetAction()
	this.Expressions = that.GetExpressions()
	this.When = that.GetWhen()
	this.RuntimeAssets = that.GetRuntimeAssets()
	return this
}

func (m *EventFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFilter) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintFilter(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Action) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFilter(dAtA, i, uint64(len(m.Action)))
		i += copy(dAtA[i:], m.Action)
	}
	if len(m.Expressions) > 0 {
		for _, s := range m.Expressions {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.When != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintFilter(dAtA, i, uint64(m.When.Size()))
		n2, err := m.When.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.RuntimeAssets) > 0 {
		for _, s := range m.RuntimeAssets {
			dAtA[i] = 0x42
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintFilter(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedEventFilter(r randyFilter, easy bool) *EventFilter {
	this := &EventFilter{}
	v1 := NewPopulatedObjectMeta(r, easy)
	this.ObjectMeta = *v1
	this.Action = string(randStringFilter(r))
	v2 := r.Intn(10)
	this.Expressions = make([]string, v2)
	for i := 0; i < v2; i++ {
		this.Expressions[i] = string(randStringFilter(r))
	}
	if r.Intn(10) != 0 {
		this.When = NewPopulatedTimeWindowWhen(r, easy)
	}
	v3 := r.Intn(10)
	this.RuntimeAssets = make([]string, v3)
	for i := 0; i < v3; i++ {
		this.RuntimeAssets[i] = string(randStringFilter(r))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedFilter(r, 9)
	}
	return this
}

type randyFilter interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneFilter(r randyFilter) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringFilter(r randyFilter) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneFilter(r)
	}
	return string(tmps)
}
func randUnrecognizedFilter(r randyFilter, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldFilter(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldFilter(dAtA []byte, r randyFilter, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateFilter(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateFilter(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateFilter(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateFilter(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateFilter(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateFilter(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateFilter(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *EventFilter) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovFilter(uint64(l))
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovFilter(uint64(l))
	}
	if len(m.Expressions) > 0 {
		for _, s := range m.Expressions {
			l = len(s)
			n += 1 + l + sovFilter(uint64(l))
		}
	}
	if m.When != nil {
		l = m.When.Size()
		n += 1 + l + sovFilter(uint64(l))
	}
	if len(m.RuntimeAssets) > 0 {
		for _, s := range m.RuntimeAssets {
			l = len(s)
			n += 1 + l + sovFilter(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFilter(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFilter(x uint64) (n int) {
	return sovFilter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventFilter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFilter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expressions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Expressions = append(m.Expressions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field When", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.When == nil {
				m.When = &TimeWindowWhen{}
			}
			if err := m.When.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuntimeAssets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RuntimeAssets = append(m.RuntimeAssets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFilter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFilter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFilter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFilter
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFilter
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFilter
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFilter(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFilter = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFilter   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("filter.proto", fileDescriptor_filter_fa80d2fd62e20657) }

var fileDescriptor_filter_fa80d2fd62e20657 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x3f, 0x4e, 0x02, 0x41,
	0x14, 0xc6, 0x19, 0x30, 0x04, 0x06, 0xd1, 0x38, 0x85, 0x59, 0x49, 0x9c, 0xd9, 0x58, 0x51, 0xe8,
	0x12, 0xb0, 0xd2, 0x4e, 0x12, 0xed, 0x8c, 0xc9, 0x46, 0x43, 0x62, 0x63, 0x76, 0x97, 0x07, 0x8c,
	0x71, 0x67, 0xc8, 0xce, 0x2c, 0xe8, 0x0d, 0x3c, 0x82, 0x25, 0x9d, 0x1c, 0xc1, 0x23, 0x50, 0x72,
	0x82, 0x8d, 0xae, 0x1d, 0x27, 0xb0, 0x34, 0x0e, 0xab, 0x41, 0xba, 0xf7, 0x7b, 0x7f, 0xe6, 0xfb,
	0xe6, 0xc3, 0x9b, 0x3d, 0xfe, 0xa0, 0x21, 0x72, 0x86, 0x91, 0xd4, 0x92, 0x54, 0x15, 0x08, 0x15,
	0x3b, 0x81, 0x8c, 0xc0, 0x19, 0xb5, 0x6a, 0x47, 0x7d, 0xae, 0x07, 0xb1, 0xef, 0x04, 0x32, 0x6c,
	0xf4, 0x65, 0x5f, 0x36, 0xcc, 0x96, 0x1f, 0xf7, 0x0c, 0x19, 0x30, 0xd5, 0xf2, 0xba, 0xb6, 0xa3,
	0x79, 0x08, 0x77, 0x63, 0x2e, 0xba, 0x72, 0x9c, 0xb5, 0x70, 0x08, 0xda, 0x5b, 0xd6, 0x07, 0xaf,
	0x79, 0x5c, 0x39, 0x1f, 0x81, 0xd0, 0x17, 0x46, 0x92, 0xdc, 0xe0, 0xd2, 0xcf, 0xb4, 0xeb, 0x69,
	0xcf, 0x42, 0x36, 0xaa, 0x57, 0x5a, 0x7b, 0xce, 0x3f, 0x7d, 0xe7, 0xca, 0xbf, 0x87, 0x40, 0x5f,
	0x82, 0xf6, 0xda, 0x74, 0x96, 0xb0, 0xdc, 0x3c, 0x61, 0x68, 0x91, 0x30, 0xf2, 0x7b, 0x76, 0x28,
	0x43, 0xae, 0x21, 0x1c, 0xea, 0x27, 0xf7, 0xef, 0x29, 0xb2, 0x8b, 0x8b, 0x5e, 0xa0, 0xb9, 0x14,
	0x56, 0xde, 0x46, 0xf5, 0xb2, 0x9b, 0x11, 0x69, 0xe2, 0x0a, 0x3c, 0x0e, 0x23, 0x50, 0x8a, 0x4b,
	0xa1, 0xac, 0x82, 0x5d, 0xa8, 0x97, 0xdb, 0xdb, 0x8b, 0x84, 0xad, 0xb6, 0xdd, 0x55, 0x20, 0x4d,
	0xbc, 0x31, 0x1e, 0x80, 0xb0, 0x8a, 0xc6, 0xdd, 0xfe, 0x9a, 0xbb, 0x6b, 0x1e, 0x42, 0xc7, 0x7c,
	0xb6, 0x33, 0x00, 0xe1, 0x9a, 0x55, 0x72, 0x82, 0xb7, 0xa2, 0x58, 0x98, 0x20, 0x3c, 0xa5, 0x40,
	0x2b, 0xab, 0x64, 0x84, 0xc8, 0x22, 0x61, 0x6b, 0x13, 0xb7, 0x9a, 0xf1, 0x99, 0xc1, 0xd3, 0xd2,
	0xf3, 0x84, 0xe5, 0xa6, 0x13, 0x86, 0xda, 0xf6, 0xd7, 0x07, 0x45, 0xd3, 0x94, 0xa2, 0xb7, 0x94,
	0xa2, 0x59, 0x4a, 0xd1, 0x3c, 0xa5, 0xe8, 0x3d, 0xa5, 0xe8, 0xe5, 0x93, 0xe6, 0x6e, 0xf3, 0xa3,
	0x96, 0x5f, 0x34, 0x91, 0x1e, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x30, 0xf8, 0x1a, 0xcc, 0xbf,
	0x01, 0x00, 0x00,
}

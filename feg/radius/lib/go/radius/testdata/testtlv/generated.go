// Code generated by radius-dict-gen. DO NOT EDIT.

package testtlv

import (
	"fbc/lib/go/radius"
	"fbc/lib/go/radius/rfc2865"
)

const (
	_TestVendor_VendorID = 9999
)

func _TestVendor_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_TestVendor_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return nil
}

func _TestVendor_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _TestVendor_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				values = append(values, vsa[2:int(vsaLen)])
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return
}

func _TestVendor_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _TestVendor_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				return vsa[2:int(vsaLen)], true
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return nil, false
}

func _TestVendor_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _TestVendor_VendorID {
			i++
			continue
		}
		for j := 0; len(vsa[j:]) >= 3; {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa[j:]) || vsaLen < 3 {
				i++
				break
			}
			if vsaTyp == typ {
				vsa = append(vsa[:j], vsa[j+int(vsaLen):]...)
			}
			j += int(vsaLen)
		}
		if len(vsa) > 0 {
			copy(p.Attributes[rfc2865.VendorSpecific_Type][i][4:], vsa)
			i++
		} else {
			p.Attributes[rfc2865.VendorSpecific_Type] = append(p.Attributes[rfc2865.VendorSpecific_Type][:i], p.Attributes[rfc2865.VendorSpecific_Type][i+i:]...)
		}
	}
	return _TestVendor_AddVendor(p, typ, attr)
}

type TestStruct struct {
	Field1 string
	Field2 uint64
}

const (
	TestStruct_Field1_Type radius.Type = 1
	TestStruct_Field2_Type radius.Type = 2
)

func newTestStruct(value TestStruct) (attribute radius.Attribute, err error) {
	var typedAttributes []radius.TypedAttribute
	var a radius.Attribute

	a, err = radius.NewString(value.Field1)
	if err != nil {
		return nil, err
	}
	typedAttributes = append(typedAttributes, radius.TypedAttribute{TestStruct_Field1_Type, a})

	a = radius.NewInteger64(value.Field2)
	if err != nil {
		return nil, err
	}
	typedAttributes = append(typedAttributes, radius.TypedAttribute{TestStruct_Field2_Type, a})

	attribute, err = radius.NewTLV(typedAttributes)
	if err != nil {
		return nil, err
	}
	return
}

func setTestStruct(a radius.Attribute) (values []TestStruct, err error) {
	var attributes radius.Attributes
	valuesLen := -1

	attributes, err = radius.TLV(a)
	if err != nil {
		return
	}

	if val, ok := attributes[TestStruct_Field1_Type]; ok {
		valuesLen = len(val)
		values = make([]TestStruct, valuesLen)
		if len(val) != valuesLen {
			err = radius.ErrTLVAttribute
		} else {
			for i := range val {
				values[i].Field1 = radius.String(val[i])
			}
		}
	} else {
		err = radius.ErrTLVAttribute
	}
	if err != nil {
		return
	}

	if val, ok := attributes[TestStruct_Field2_Type]; ok {
		if len(val) != valuesLen {
			err = radius.ErrTLVAttribute
		} else {
			for i := range val {
				values[i].Field2, err = radius.Integer64(val[i])
			}
		}
	} else {
		err = radius.ErrTLVAttribute
	}
	if err != nil {
		return
	}
	return
}

func TestStruct_Add(p *radius.Packet, values []TestStruct) error {
	var attribute radius.Attribute
	if len(values) < 1 {
		return radius.ErrEmptyStruct
	}
	for _, value := range values {
		_attr, _err := newTestStruct(value)
		if _err != nil {
			return _err
		}
		attribute = append(attribute, _attr...)
	}
	return _TestVendor_AddVendor(p, 4, attribute)
}

func TestStruct_Get(p *radius.Packet) (value []TestStruct) {
	value, _ = TestStruct_Lookup(p)
	return
}

func TestStruct_Gets(p *radius.Packet) (values [][]TestStruct, err error) {
	var value []TestStruct
	for _, attr := range _TestVendor_GetsVendor(p, 4) {
		value, err = setTestStruct(attr)
		if err != nil {
			return
		}
		values = append(values, []TestStruct(value))
	}
	return
}

func TestStruct_Lookup(p *radius.Packet) (values []TestStruct, err error) {
	a, ok := _TestVendor_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	values, err = setTestStruct(a)
	return
}

func TestStruct_Set(p *radius.Packet, values []TestStruct) error {
	var attribute radius.Attribute
	for _, value := range values {
		_attr, _err := newTestStruct(value)
		if _err != nil {
			return _err
		}
		attribute = append(attribute, _attr...)
	}
	return _TestVendor_SetVendor(p, 4, attribute)
}

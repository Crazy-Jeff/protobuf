// Code generated by protoc-gen-gogo.
// source: extension.proto
// DO NOT EDIT!

package extension

import testing "testing"
import math_rand "math/rand"
import time "time"
import code_google_com_p_gogoprotobuf_proto "code.google.com/p/gogoprotobuf/proto"
import testing1 "testing"
import math_rand1 "math/rand"
import time1 "time"
import encoding_json "encoding/json"
import testing2 "testing"
import math_rand2 "math/rand"
import time2 "time"
import code_google_com_p_gogoprotobuf_proto1 "code.google.com/p/gogoprotobuf/proto"
import math_rand3 "math/rand"
import time3 "time"
import testing3 "testing"
import fmt "fmt"
import math_rand4 "math/rand"
import time4 "time"
import testing4 "testing"
import code_google_com_p_gogoprotobuf_proto2 "code.google.com/p/gogoprotobuf/proto"
import math_rand5 "math/rand"
import time5 "time"
import testing5 "testing"
import fmt1 "fmt"
import go_parser "go/parser"
import math_rand6 "math/rand"
import time6 "time"
import testing6 "testing"
import code_google_com_p_gogoprotobuf_proto3 "code.google.com/p/gogoprotobuf/proto"

func TestMyExtendableProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedMyExtendable(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &MyExtendable{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOtherExtenableProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedOtherExtenable(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &OtherExtenable{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestMyExtendableJSON(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedMyExtendable(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &MyExtendable{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestOtherExtenableJSON(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedOtherExtenable(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &OtherExtenable{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestMyExtendableProtoText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedMyExtendable(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &MyExtendable{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestMyExtendableProtoCompactText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedMyExtendable(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	msg := &MyExtendable{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOtherExtenableProtoText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedOtherExtenable(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &OtherExtenable{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOtherExtenableProtoCompactText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedOtherExtenable(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	msg := &OtherExtenable{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestMyExtendableStringer(t *testing3.T) {
	popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
	p := NewPopulatedMyExtendable(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestOtherExtenableStringer(t *testing3.T) {
	popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
	p := NewPopulatedOtherExtenable(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestMyExtendableSize(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedMyExtendable(popr, true)
	data, err := code_google_com_p_gogoprotobuf_proto2.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
}

func TestOtherExtenableSize(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedOtherExtenable(popr, true)
	data, err := code_google_com_p_gogoprotobuf_proto2.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
}

func TestMyExtendableGoString(t *testing5.T) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	p := NewPopulatedMyExtendable(popr, false)
	s1 := p.GoString()
	s2 := fmt1.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestOtherExtenableGoString(t *testing5.T) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	p := NewPopulatedOtherExtenable(popr, false)
	s1 := p.GoString()
	s2 := fmt1.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestMyExtendableVerboseEqual(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedMyExtendable(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto3.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &MyExtendable{}
	if err := code_google_com_p_gogoprotobuf_proto3.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestOtherExtenableVerboseEqual(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedOtherExtenable(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto3.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &OtherExtenable{}
	if err := code_google_com_p_gogoprotobuf_proto3.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}

//These tests are generated by code.google.com/p/gogoprotobuf/plugin/testgen

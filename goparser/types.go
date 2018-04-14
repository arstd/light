package goparser

import (
	"fmt"
	"go/types"
	"reflect"
	"strings"

	"github.com/arstd/log"
)

type Store struct {
	Source string
	Log    bool

	Package string            // store
	Imports map[string]string // database/sql => sql
	Name    string            // User

	Methods []*Method

	Init func() string
}

func (store *Store) Initx() string {
	name := store.Name
	if name[0] == 'I' {
		name = name[1:]
		return fmt.Sprintf("func init(){ %s = new(Store%s)}\ntype Store%s struct{}\n", name, name, name)
	} else {
		return fmt.Sprintf("func init(){ Default%s = new(Store%s)}\ntype Store%s struct{}\n", name, name, name)
	}
}

func (s *Store) MethodByName(name string) *Method {
	for _, a := range s.Methods {
		if a.Name == name {
			return a
		}
	}
	return nil
}

type Method struct {
	Store *Store `json:"-"`

	Name string // Insert
	Doc  string // insert into users ...

	Params  *Tuple
	Results *Tuple

	ResultTypeName     func() string
	ResultTypeWrap     func() string
	ResultElemTypeName func() string
	ResultVarByTagScan func(name string) string
	ParamsVarByName    func(string) *Var
	Tx                 func() string
}

func NewMethod(store *Store, name, doc string) *Method {
	m := &Method{Store: store, Name: name, Doc: doc}
	m.ResultTypeName = func() string { return m.Results.Result().TypeName() }
	m.ResultTypeWrap = func() string { return m.Results.Result().Wrap(true) }
	m.ResultElemTypeName = func() string { return m.Results.Result().ElemTypeName() }
	m.ResultVarByTagScan = func(name string) string {
		s := m.Results.Result()
		v := s.VarByTag(name)
		return v.Scan("xu." + v.VName)
	}

	m.Tx = func() string {
		for i := 0; i < m.Params.Len(); i++ {
			v := m.Params.At(i)
			typ := typeString(v.Store, v.Type())
			if typ == "*sql.Tx" {
				return v.Name()
			}
		}
		return ""
	}
	return m
}

func (m *Method) Signature() string {
	name := m.Store.Name
	if name[0] == 'I' {
		name = name[1:]
	}
	return "func (*Store" + name + ")" + m.Name +
		"(" + m.Params.String() + ")(" + m.Results.String() + "){\n"
}

type Tuple struct {
	Store *Store `json:"-"`
	*types.Tuple
}

func (t *Tuple) String() string {
	var ss []string
	for i := 0; i < t.Len(); i++ {
		ss = append(ss, t.At(i).String())
	}
	return strings.Join(ss, ", ")
}

func (t *Tuple) At(i int) *Var {
	x := t.Tuple.At(i)
	return &Var{VName: x.Name(), Store: t.Store, Var: t.Tuple.At(i)}
}

func (t *Tuple) VarByName(name string) *Var {
	name = strings.Trim(name, "`")
	if name == "" {
		panic("name must not blank")
	}

	var v *Var
	parts := strings.Split(name, ".")

	parts0 := lowerCamelCase(parts[0])
	// 从参数列表中查找
	for i := 0; i < t.Len(); i++ {
		x := t.At(i)
		if x.Name() == parts0 {
			v = x
			break
		}
	}
	// 如果找到了
	if v != nil {
		switch len(parts) {
		case 1:
			return v

		case 2:
			s := underlying(v.Type())
			for i := 0; i < s.NumFields(); i++ {
				x := s.Field(i)
				if x.Name() == parts[1] {
					z := getTag(s.Tag(i), "light")
					return &Var{VName: name, Store: t.Store, Var: x, Tag: z}
				}
			}
			panic("variable " + name + " not exist")

		default:
			panic("variable " + name + " to long")
		}
	}

	// 从结构体参数中查找
	if len(parts) > 1 {
		panic("variable " + name + " not exist")
	}

	out := upperCamelCase(name)
	for i := 0; i < t.Len(); i++ {
		s := underlying(t.At(i).Type())
		if s != nil {
			for j := 0; j < s.NumFields(); j++ {
				x := s.Field(j)
				if x.Name() == out {
					z := getTag(s.Tag(j), "light")
					return &Var{
						VName: t.At(i).Name() + "." + x.Name(),
						Store: t.Store,
						Var:   x,
						Tag:   z,
					}
				}
			}
		}
	}
	panic("variable " + name + " not exist")
}

func (t *Tuple) Result() *Var {
	switch t.Len() {
	case 1:
		panic("unimplemented")
	case 2:
		return t.At(0)
	case 3:
		return t.At(1)
	default:
		panic(t.Len())
	}
}

type Var struct {
	VName string
	Store *Store `json:"-"`
	Tag   string
	*types.Var
}

func (v *Var) VarByTag(field string) *Var {
	field = strings.Trim(field, "`")
	s := underlying(v.Type())
	for i := 0; i < s.NumFields(); i++ {
		tag := s.Tag(i)
		t := getTag(tag, "light")
		if t != "" {
			tt := strings.Split(t, ",")
			if tt[0] != "" {
				if strings.HasPrefix(t, tt[0]) {
					return &Var{
						VName: s.Field(i).Name(),
						Store: v.Store,
						Var:   s.Field(i),
						Tag:   t,
					}
				}
			}
		}
	}

	out := upperCamelCase(field)
	for i := 0; i < s.NumFields(); i++ {
		x := s.Field(i)
		if strings.EqualFold(out, x.Name()) {
			t := getTag(s.Tag(i), "light")
			return &Var{
				VName: s.Field(i).Name(),
				Store: v.Store,
				Var:   s.Field(i),
				Tag:   t,
			}
		}
	}
	panic(field + " not found")
}

func lowerCamelCase(field string) (out string) {
	return camelCase(field, false)
}

func upperCamelCase(field string) (out string) {
	return camelCase(field, true)
}

func camelCase(name string, first bool) (out string) {
	for _, v := range name {
		if first {
			out += strings.ToUpper(string(v))
			first = false
		} else if v == '_' {
			first = true
		} else {
			out += string(v)
			first = false
		}
	}
	return out
}

func getTag(tag, key string) string {
	idx := strings.Index(tag, key+`:"`)
	if idx == -1 {
		return ""
	}
	tag = tag[idx+len(key)+2:]
	idx = strings.Index(tag, `"`)
	if idx == -1 {
		panic(tag)
	}
	return tag[:idx]
}

func (v *Var) IsBasic() bool {
	_, ok := v.Type().(*types.Basic)
	return ok
}

func (v *Var) NotDefault(name string) string {
	switch u := v.Type().(type) {
	case *types.Named:
		if u.String() == "time.Time" {
			return "!" + name + ".IsZero()"
		}

		t, ok := u.Underlying().(*types.Basic)
		if !ok {
			log.Fatalf("%#v", u)
		}
		bi := t.Info()
		switch {
		case bi&types.IsString == types.IsString:
			return name + ` != ""`
		case bi&types.IsInteger == types.IsInteger:
			return name + ` != 0`
		case bi&types.IsFloat == types.IsFloat:
			return name + ` != 0`
		case bi&types.IsBoolean == types.IsBoolean:
			return name
		default:
			panic(t)
		}

	case *types.Basic:
		bi := u.Info()
		switch {
		case bi&types.IsString == types.IsString:
			return name + ` != ""`
		case bi&types.IsInteger == types.IsInteger:
			return name + ` != 0`
		case bi&types.IsFloat == types.IsFloat:
			return name + ` != 0`
		case bi&types.IsBoolean == types.IsBoolean:
			return name
		default:
			panic(u)
		}

	case *types.Pointer:
		return name + " != nil"

	case *types.Struct:
		return name + " != nil"

	case *types.Slice:
		return "len(" + name + ") != 0"

	default:
		panic(" unimplement " + reflect.TypeOf(u).String() + u.String())
	}
}

func (v *Var) Value(name string) string {
	if v.Wrap() == "" {
		return name
	}
	return v.Wrap() + "(&" + name + ")"
}

func (v *Var) Scan(name string) string {
	s := v.Value(name)
	if strings.HasPrefix(s, "null.") {
		return s
	}
	return "&" + s
}

func (v *Var) Nullable() bool {
	for i, v := range strings.Split(v.Tag, ",") {
		if i == 0 {
			continue
		}
		if v == "nullable" {
			return true
		}
	}
	return false
}

func (v *Var) IsSlice() bool {
	_, ok := v.Type().(*types.Slice)
	return ok
}

func (v *Var) Wrap(force ...bool) string {
	switch u := v.Type().(type) {
	case *types.Pointer, *types.Named, *types.Slice, *types.Array:
		return ""

	case *types.Basic:
		if v.Nullable() || (len(force) > 0 && force[0]) {
			name := u.Name()
			return "null." + strings.ToUpper(name[:1]) + name[1:]
		}
		return ""

	default:
		panic(reflect.TypeOf(u))
	}
}

func underlying(t types.Type) *types.Struct {
	switch u := t.(type) {
	case *types.Named:
		return underlying(u.Underlying())

	case *types.Pointer:
		return underlying(u.Elem())

	case *types.Slice:
		return underlying(u.Elem())

	case *types.Struct:
		return u

	default:
		return nil
	}
}

func (v *Var) String() string {
	typ := typeString(v.Store, v.Type())
	name := v.Name()
	return name + " " + typ
}

func (v *Var) TypeName() string {
	return strings.TrimLeft(v.String(), " *")
}

func (v *Var) ElemTypeName() string {
	return strings.TrimLeft(v.String(), " []*")
}

func typeString(store *Store, t types.Type) string {
	switch u := t.(type) {
	case *types.Named:
		if obj := u.Obj(); obj != nil {
			if pkg := obj.Pkg(); pkg != nil {
				path := pkg.Path()
				if path != "" && path[0] != '/' {
					store.Imports[pkg.Path()] = ""
					return shortPkg(pkg.Path()) + "." + obj.Name()
				}
			}
			return obj.Name()
		}
		return typeString(store, u.Underlying())

	case *types.Basic:
		return u.String()

	case *types.Pointer:
		return "*" + typeString(store, u.Elem())

	case *types.Struct:
		return u.String()

	case *types.Slice:
		return "[]" + typeString(store, u.Elem())

	default:
		panic(" unimplement " + reflect.TypeOf(u).String())
	}
}

func shortPkg(path string) string {
	return path[strings.LastIndex(path, "/")+1:]
}

package interfaces

type InterfaceBuilder struct {
	interfaces []interface{}
}

func NewInterfaceBuilder() *InterfaceBuilder {
	var interfaces InterfaceBuilder
	return &interfaces
}

func (builder *InterfaceBuilder) Append(arg interface{}) *InterfaceBuilder {
	builder.interfaces = append(builder.interfaces, arg)
	return builder
}

func (builder *InterfaceBuilder) Appends(args ...interface{}) *InterfaceBuilder {
	for i := range args {
		builder.interfaces = append(builder.interfaces, args[i])
	}
	return builder
}

func (builder *InterfaceBuilder) Clear() *InterfaceBuilder {
	var interfaces []interface{}
	builder.interfaces = interfaces
	return builder
}

func (builder *InterfaceBuilder) ToInterfaces() []interface{} {
	return builder.interfaces
}

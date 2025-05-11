package main

import (
	"fmt"
	"runtime"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/GDExtension"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/startup"
	"graphics.gd/variant"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

/*
HelloWorld is a simple extension to demonstrate how to export
Go methods so that they can be used in scripts.
*/
type HelloWorld struct {
	Object.Extension[HelloWorld]
}

func NewHelloWorld() *HelloWorld {
	fmt.Println("NewHelloWorld!")
	return &HelloWorld{}
}

func Static() string {
	return "static method call"
}

// Print prints "Hello World"
func (h *HelloWorld) Print() { fmt.Println("Hello World") }

// Echo prints the given string, signalling that it
// was printed by Go code.
func (h *HelloWorld) Echo(s string) { fmt.Println(s + " from Go!") }

// Arch returns the current GOARCH value.
func (h *HelloWorld) Arch() string { return runtime.GOARCH }

func (h *HelloWorld) Bar(message string) *Bar {
	return &Bar{
		Message: message,
	}
}

type Rotator struct {
	Sprite2D.Extension[Rotator]

	Map map[string]int
}

func (r *Rotator) Process(delta Float.X) {
	node2D := r.AsNode2D()
	node2D.SetRotation(node2D.Rotation() + delta)
}

type StartedSignalEmitter struct {
	Node.Extension[StartedSignalEmitter]

	started chan<- struct{}
}

func (r *StartedSignalEmitter) Ready() {
	select {
	case r.started <- struct{}{}:
	default:
	}
}

/*
ExtendedNode demonstrates how to call the methods of builtin objects.
*/
type ExtendedNode struct {
	Node2D.Extension[ExtendedNode]

	StringField string
}

func (e *ExtendedNode) Ready() {
	fmt.Println("Ready!")

	node := e.AsNode2D()

	fmt.Println("class:", Object.Instance(node.AsObject()).String())

	var obj = Object.New()
	fmt.Println(obj.ClassName())

	fmt.Println(Engine.GetSingletonList())
	fmt.Println("Scene is ready!")

	fmt.Println("sin=", Angle.Sin(1.5))

	fmt.Println("rotation=", node.Rotation())
	node.SetRotation(3.14)
	fmt.Println("rotation=", node.Rotation())

	pos := node.Position()

	fmt.Println("position=", pos)

	pos.X = 100

	node.SetPosition(pos)
	fmt.Println("position=", pos)

	v := variant.New(node)
	fmt.Printf("result: %[1]T %[1]v\n", v.Call("Position"))
}

type Bar struct {
	Object.Extension[Bar]

	Message string
}

// main init function, where the extensions are exported so that
// they are available to the engine.
func main() {
	classdb.Register[HelloWorld](NewHelloWorld, Static, map[string]any{
		"get_bar": (*HelloWorld).Bar,
	})
	classdb.Register[ExtendedNode]()
	classdb.Register[Rotator]()
	classdb.Register[StartedSignalEmitter]()
	classdb.Register[Bar]()

	startup.LoadingScene()
	fmt.Println("Engine Version is: ", Engine.Version())
	fmt.Println("Extension: ", GDExtension.LibraryPath())
	startup.Scene()
	fmt.Println("Shutting Down...")
}

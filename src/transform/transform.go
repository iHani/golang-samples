package transform

type Transform struct {
	Name string
}

type Tchild struct {
	Transform //an anonymous field of type Transform
}

func (t Transform) Doit() string {
   return t.Name + " is it!"
}

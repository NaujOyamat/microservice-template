package applications

type GreettingApplication struct{}

func (g *GreettingApplication) HelloWorld() string {
	return "Hola Mundo...!"
}

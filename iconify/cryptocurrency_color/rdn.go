package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rdn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#2A2A2A"/><path fill="#FFF" d="M10 6.006h8.111c.098-.03.148.06.212.11a14.88 14.88 0 0 1 2.046 2.39a15.153 15.153 0 0 1 1.972 3.946c.482 1.484.71 3.054.65 4.614h-4.135a5 5 0 0 0 .022-.624c-.043-1.471-.497-2.914-1.212-4.192c-.765-1.369-1.833-2.55-3.047-3.531c-1.322-1.069-2.81-1.913-4.36-2.595c-.086-.04-.176-.072-.259-.118zm3.002 7.327a4715.11 4715.11 0 0 1 4.155 7.533c.333.606.67 1.21 1 1.817C16.436 23.785 14.72 24.897 13 26V13.333z"/></g>`),
		g.Group(children),
	)
}
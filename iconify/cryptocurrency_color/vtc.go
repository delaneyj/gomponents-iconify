package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vtc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#048657"/><path fill="#FFF" fill-rule="nonzero" d="m2.632 16.57l1.95-2.221h6.938l4.263 5.438L26.938 4.334a14.205 14.205 0 0 1 1.86 2.04a15.002 15.002 0 0 1 1.496 2.446L16.599 27.763c-.259.272-.531.408-.816.408c-.286 0-.573-.136-.862-.408L6.26 16.569H2.632z"/></g>`),
		g.Group(children),
	)
}
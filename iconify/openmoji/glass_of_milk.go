package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassOfMilk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FFF" d="M53.545 27.939L48.35 67H24.6l-5.921-38.951"/><path fill="#D0CFCE" d="M43.404 29L38.35 67h10l5.054-38z"/><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M48.35 67H24.6L17 17h38z"/><path d="M19 28h35"/></g>`),
		g.Group(children),
	)
}
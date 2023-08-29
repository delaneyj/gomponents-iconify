package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpressionlessFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FCEA2B" d="M36 13c-12.682 0-23 10.318-23 23s10.318 23 23 23s23-10.318 23-23s-10.318-23-23-23z"/><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><circle cx="36" cy="36" r="23"/><path stroke-linecap="round" stroke-linejoin="round" d="M27 43h18M25 31h5m13 0h5"/></g>`),
		g.Group(children),
	)
}
package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Multiply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D0CFCE" d="M30 36L18 24l6-6l12 12l12-12l6 6l-12 12l12 12l-6 6l-12-12l-12 12l-6-6z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M30 36L18 24l6-6l12 12l12-12l6 6l-12 12l12 12l-6 6l-12-12l-12 12l-6-6z"/>`),
		g.Group(children),
	)
}
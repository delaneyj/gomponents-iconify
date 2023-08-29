package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRightFilledOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.52 10.136a1 1 0 0 1 0 1.728l-9.016 5.259A1 1 0 0 1 7 16.259V5.74a1 1 0 0 1 1.504-.864l9.015 5.26Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M16 10a.5.5 0 0 1-.243.429l-10 6A.5.5 0 0 1 5 16V4a.5.5 0 0 1 .757-.429l10 6A.5.5 0 0 1 16 10ZM6 4.883v10.234L14.528 10L6 4.883Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}
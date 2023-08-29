package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleUpFilledOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M10.636 5.98a1 1 0 0 1 1.728 0l5.259 9.016a1 1 0 0 1-.864 1.504H6.24a1 1 0 0 1-.864-1.504l5.26-9.015Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M10 4a.5.5 0 0 1 .429.243l6 10A.5.5 0 0 1 16 15H4a.5.5 0 0 1-.429-.757l6-10A.5.5 0 0 1 10 4ZM4.883 14h10.234L10 5.472L4.883 14Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}
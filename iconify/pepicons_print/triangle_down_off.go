package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDownOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M11.75 14.014L15.528 7.5H7.972l3.778 6.514Zm-.865 2.495a1 1 0 0 0 1.73 0l5.514-9.507a1 1 0 0 0-.865-1.502H6.236a1 1 0 0 0-.865 1.502l5.514 9.507Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M10 16a.5.5 0 0 1-.429-.243l-6-10A.5.5 0 0 1 4 5h12a.5.5 0 0 1 .429.757l-6 10A.5.5 0 0 1 10 16Zm5.117-10H4.883L10 14.528L15.117 6Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}
package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleUpFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.636 5.98a1 1 0 0 1 1.728 0l5.259 9.016a1 1 0 0 1-.864 1.504H6.24a1 1 0 0 1-.864-1.504l5.26-9.015Z" opacity=".2"/><path d="M10 4a.5.5 0 0 1 .429.243l6 10A.5.5 0 0 1 16 15H4a.5.5 0 0 1-.429-.757l6-10A.5.5 0 0 1 10 4ZM4.883 14h10.234L10 5.472L4.883 14Z"/></g>`),
		g.Group(children),
	)
}
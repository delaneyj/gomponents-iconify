package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.75 8.486L6.972 15h7.556L10.75 8.486Zm.865-2.495a1 1 0 0 0-1.73 0l-5.514 9.507A1 1 0 0 0 5.236 17h11.028a1 1 0 0 0 .865-1.502l-5.514-9.507Z" opacity=".2"/><path d="M10 4a.5.5 0 0 1 .429.243l6 10A.5.5 0 0 1 16 15H4a.5.5 0 0 1-.429-.757l6-10A.5.5 0 0 1 10 4ZM4.883 14h10.234L10 5.472L4.883 14Z"/></g>`),
		g.Group(children),
	)
}
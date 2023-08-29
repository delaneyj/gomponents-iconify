package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberNineSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-14Z"/><path fill="currentColor" d="M9.955 16.315a1.25 1.25 0 1 0 2.09 1.37l-2.09-1.37Zm5.615-4.01a1.25 1.25 0 1 0-2.09-1.37l2.09 1.37ZM10.25 10c0-.966.784-1.75 1.75-1.75v-2.5A4.25 4.25 0 0 0 7.75 10h2.5ZM12 8.25c.966 0 1.75.784 1.75 1.75h2.5A4.25 4.25 0 0 0 12 5.75v2.5ZM13.75 10A1.75 1.75 0 0 1 12 11.75v2.5A4.25 4.25 0 0 0 16.25 10h-2.5ZM12 11.75A1.75 1.75 0 0 1 10.25 10h-2.5A4.25 4.25 0 0 0 12 14.25v-2.5Zm.046 5.935l3.524-5.38l-2.09-1.37l-3.525 5.38l2.09 1.37Z"/></g>`),
		g.Group(children),
	)
}
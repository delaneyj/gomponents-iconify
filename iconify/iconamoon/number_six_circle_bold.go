package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSixCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"/><path fill="currentColor" d="M14.046 7.685a1.25 1.25 0 0 0-2.092-1.37l2.091 1.37Zm-5.617 4.01a1.25 1.25 0 1 0 2.091 1.37l-2.09-1.37ZM13.75 14A1.75 1.75 0 0 1 12 15.75v2.5A4.25 4.25 0 0 0 16.25 14h-2.5ZM12 15.75A1.75 1.75 0 0 1 10.25 14h-2.5A4.25 4.25 0 0 0 12 18.25v-2.5ZM10.25 14c0-.966.784-1.75 1.75-1.75v-2.5A4.25 4.25 0 0 0 7.75 14h2.5ZM12 12.25c.966 0 1.75.784 1.75 1.75h2.5A4.25 4.25 0 0 0 12 9.75v2.5Zm-.046-5.935l-3.525 5.38l2.091 1.37l3.526-5.38l-2.092-1.37Z"/></g>`),
		g.Group(children),
	)
}
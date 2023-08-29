package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 13l4-4l-4-4M7 13L3 9l4-4"/><path d="M12 14a5 5 0 0 1 5-5h4m-9 10v-5a5 5 0 0 0-5-5H3"/></g>`),
		g.Group(children),
	)
}
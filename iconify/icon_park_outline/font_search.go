package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><circle cx="22.834" cy="22.834" r="17" stroke-linejoin="round"/><path d="m35 35l6 6"/><path stroke-linejoin="round" d="M23 17v14m-5-14h10"/></g>`),
		g.Group(children),
	)
}
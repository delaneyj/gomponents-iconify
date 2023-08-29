package heroicons_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 3.055A9.001 9.001 0 1 0 20.945 13H11V3.055Z"/><path d="M20.488 9H15V3.512A9.025 9.025 0 0 1 20.488 9Z"/></g>`),
		g.Group(children),
	)
}
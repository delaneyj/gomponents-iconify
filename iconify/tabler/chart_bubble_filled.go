package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBubbleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M6 12a4 4 0 1 1-3.995 4.2L2 16l.005-.2A4 4 0 0 1 6 12zm10 4a3 3 0 1 1-2.995 3.176L13 19l.005-.176A3 3 0 0 1 16 16zM14.5 2a5.5 5.5 0 1 1-5.496 5.721L9 7.5l.004-.221A5.5 5.5 0 0 1 14.5 2z"/></g>`),
		g.Group(children),
	)
}
package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DashboardReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4zm20 1h2v5h-2zm-4-5h2v10h-2zm-4 2h2v8h-2z"/><path fill="currentColor" d="M28 2H4a2.002 2.002 0 0 0-2 2v12h2v-3h24.001l.001 15H16v2h12a2.003 2.003 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2Zm-16 9H4V4h8Zm2 0V4h14v7Z"/>`),
		g.Group(children),
	)
}
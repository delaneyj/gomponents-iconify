package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartClusterBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 30H4a2 2 0 0 1-2-2V2h2v26h26Z"/><path fill="currentColor" d="M10 16h2v10h-2zm-3 6h2v4H7zM26 8h2v18h-2zm-3 6h2v12h-2zm-6 12h-2V12h2zm3 0h-2v-8h2z"/>`),
		g.Group(children),
	)
}
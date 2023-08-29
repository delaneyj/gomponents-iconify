package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitPullLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6zm0 0v10m12-6a3 3 0 1 0 0 6a3 3 0 0 0 0-6zm0 0V9a2 2 0 0 0-2-2h-1m-2 0l2-2v2m-2 0h2m-2 0l2 2V7"/>`),
		g.Group(children),
	)
}
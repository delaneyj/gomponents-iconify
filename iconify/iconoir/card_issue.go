package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardIssue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 9V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8M22 9H6m16 0v4m-1 3.05a3.5 3.5 0 1 0-5 4.9m4.998-4.9A3.5 3.5 0 1 1 16 20.95m5-4.9l-5 4.9"/>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCherryPick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12a3 3 0 1 0 6 0a3 3 0 1 0-6 0m3-9v6m0 6v6m6-14h2.5l1.5 5l-1.5 5H13m4-5h3"/>`),
		g.Group(children),
	)
}
package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DashboardLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21V11h8v10h-8ZM3 13V3h8v10H3Zm6-2V5H5v6h4ZM3 21v-6h8v6H3Zm2-2h4v-2H5v2Zm10 0h4v-6h-4v6ZM13 3h8v6h-8V3Zm2 2v2h4V5h-4Z"/>`),
		g.Group(children),
	)
}
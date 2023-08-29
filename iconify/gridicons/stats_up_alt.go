package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatsUpAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 21H3v-2h18v2zM8 10H4v7h4v-7zm6-4h-4v11h4V6zm6-3h-4v14h4V3z"/>`),
		g.Group(children),
	)
}
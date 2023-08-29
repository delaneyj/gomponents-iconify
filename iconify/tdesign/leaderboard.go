package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leaderboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 3h8v8h6v10H2V9h6V3Zm2 16h4V5h-4v14Zm6 0h4v-6h-4v6Zm-8 0v-8H4v8h4Z"/>`),
		g.Group(children),
	)
}
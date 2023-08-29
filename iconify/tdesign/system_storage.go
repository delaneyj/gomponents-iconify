package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemStorage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 4v3.625H11V4H8.5ZM13 4v5.625H6.5V4H4v16h16V9.04L14.96 4H13Zm2.79-2L22 8.212v13.79H2V2h13.79ZM7 14v-2h10v2H7Zm0 3v-2h6v2H7Z"/>`),
		g.Group(children),
	)
}
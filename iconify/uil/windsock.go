package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windsock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.08 5L10 4.33l-3-.25V3a1 1 0 0 0-2 0v18a1 1 0 0 0 2 0v-7.08l3-.25l8.08-.67a1 1 0 0 0 .92-1V6a1 1 0 0 0-.92-1ZM9 11.75l-2 .16V6.09l2 .16Zm4-.34l-2 .17V6.42l2 .17Zm4-.33l-2 .17v-4.5l2 .17Z"/>`),
		g.Group(children),
	)
}
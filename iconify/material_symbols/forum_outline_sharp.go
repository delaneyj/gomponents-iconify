package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForumOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18v-3h13V6h3v16l-4-4H6Zm-4-1V2h15v11H6l-4 4Zm13-6V4H4v7h11ZM4 11V4v7Z"/>`),
		g.Group(children),
	)
}
package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 5H4v2h12V5Zm0 4H4v2h12V9ZM4 13h8v2H4v-2Zm16 3l-6-3v6l6-3Z"/>`),
		g.Group(children),
	)
}
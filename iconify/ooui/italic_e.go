package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ItalicE(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 19L8 1h9l-.6 2H10L8.6 9H15l-.6 2H8.2L7 17h6.5l-.5 2H4Z"/>`),
		g.Group(children),
	)
}
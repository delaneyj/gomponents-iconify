package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sweden(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 4H10v7h13V4Zm0 9v7H10v-7h13ZM8 13v7H1v-7h7Zm-7-2V4h7v7H1Z"/>`),
		g.Group(children),
	)
}
package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 4h8v2H8V4Zm10 2h4v12h-4v4H6v-4H2V6h4V2h12v4Zm2 10h-2v-2H6v2H4V8h16v8ZM8 16h8v4H8v-4Zm0-6H6v2h2v-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KitchenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 8h2V5H8v3Zm0 9h2v-5H8v5Zm4-5ZM4 22V11h16v11H4ZM4 9V2h16v7H4Z"/>`),
		g.Group(children),
	)
}
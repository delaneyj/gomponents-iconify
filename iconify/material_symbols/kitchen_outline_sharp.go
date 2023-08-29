package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KitchenOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 8V5h2v3H8Zm0 9v-5h2v5H8Zm-4 5V2h16v20H4Zm2-2h12v-9H6v9ZM6 9h12V4H6v5Z"/>`),
		g.Group(children),
	)
}
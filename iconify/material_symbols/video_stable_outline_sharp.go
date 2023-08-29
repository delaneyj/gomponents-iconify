package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoStableOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.025 12Zm4.95 6l2.3-8.65L7.075 6l-2.3 8.65l12.2 3.35ZM2 20V4h20v16H2Zm2-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}
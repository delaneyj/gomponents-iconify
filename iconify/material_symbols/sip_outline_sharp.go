package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SipOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15h2V9h-2v6Zm3 0h1.5v-2H19V9h-5v6Zm-9 0h5v-3.75H6.5v-.75H10V9H5v3.75h3.5v.75H5V15Zm10.5-3.5v-1h2v1h-2ZM2 20V4h20v16H2Zm2-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}
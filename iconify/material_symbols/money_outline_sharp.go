package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 16h5V8h-5v8Zm2-2v-4h1v4h-1Zm-8 2h5V8H8v8Zm2-2v-4h1v4h-1Zm-5 2h2V8H5v8Zm-3 4V4h20v16H2ZM4 6v12V6Zm0 12h16V6H4v12Z"/>`),
		g.Group(children),
	)
}
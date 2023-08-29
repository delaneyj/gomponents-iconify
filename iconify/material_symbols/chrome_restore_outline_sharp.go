package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChromeRestoreOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 16V4h12v12Zm2-2h8V6h-8Zm-6 6V8.525h2V18h9.475v2Zm6-6V6v8Z"/>`),
		g.Group(children),
	)
}
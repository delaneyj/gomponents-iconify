package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HangoutVideoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 16h8v-3.2l4 3.2V8l-4 3.2V8H6v8Zm-4 4V4h20v16H2Z"/>`),
		g.Group(children),
	)
}
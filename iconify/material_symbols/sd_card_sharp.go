package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdCardSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 11h2V7H9v4Zm3 0h2V7h-2v4Zm3 0h2V7h-2v4ZM4 22V8l6-6h10v20H4Z"/>`),
		g.Group(children),
	)
}
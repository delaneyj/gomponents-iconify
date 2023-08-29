package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageUsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17V7h2v8h3V7h2v10H4Zm9 0v-3h2v1h3v-2h-5V7h7v3h-2V9h-3v2h5v6h-7Z"/>`),
		g.Group(children),
	)
}
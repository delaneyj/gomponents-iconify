package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountTreeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 21v-3h-4V8H9v3H2V3h7v3h6V3h7v8h-7V8h-2v8h2v-3h7v8h-7ZM4 5v4v-4Zm13 10v4v-4Zm0-10v4v-4Zm0 4h3V5h-3v4Zm0 10h3v-4h-3v4ZM4 9h3V5H4v4Z"/>`),
		g.Group(children),
	)
}
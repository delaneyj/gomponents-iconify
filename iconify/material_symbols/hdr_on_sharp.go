package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdrOnSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 15V9h5v3.9h-.9L21 15h-1.5l-.9-2h-1.1v2H16Zm1.5-3.5h2v-1h-2v1ZM3 15V9h1.5v2h2V9H8v6H6.5v-2.5h-2V15H3Zm6.5 0V9H13q.6 0 1.05.45t.45 1.05v3q0 .6-.45 1.05T13 15H9.5Zm1.5-1.5h2v-3h-2v3Z"/>`),
		g.Group(children),
	)
}
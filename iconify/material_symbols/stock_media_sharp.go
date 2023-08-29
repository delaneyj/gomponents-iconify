package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StockMediaSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 15V2h13v13H2Zm6-3.75L6.8 9.6L5 12h7L9.7 9L8 11.25ZM17 22q-1.25 0-2.125-.875T14 19q0-1.25.875-2.125T17 16q.275 0 .513.05t.487.125V11h4v2h-2v6q0 1.25-.875 2.125T17 22Z"/>`),
		g.Group(children),
	)
}
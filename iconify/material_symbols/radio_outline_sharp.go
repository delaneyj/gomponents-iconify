package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V6.65L15.9 1l.65 1.65L8.3 6H22v16H2Zm2-2h16v-7H4v7Zm4-1q1.05 0 1.775-.725T10.5 16.5q0-1.05-.725-1.775T8 14q-1.05 0-1.775.725T5.5 16.5q0 1.05.725 1.775T8 19Zm-4-8h12V9h2v2h2V8H4v3Zm0 9v-7v7Z"/>`),
		g.Group(children),
	)
}
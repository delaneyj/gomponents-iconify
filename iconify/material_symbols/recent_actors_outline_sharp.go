package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecentActorsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 19V5h14v14H1Zm2-3.65q1.1-.65 2.35-1T8 14q1.4 0 2.65.35t2.35 1V7H3v8.35ZM8 16q-1.025 0-2 .25T4.15 17h7.7q-.875-.5-1.85-.75T8 16Zm0-2.75q-1.125 0-1.938-.813T5.25 10.5q0-1.125.813-1.938T8 7.75q1.125 0 1.938.813t.812 1.937q0 1.125-.813 1.938T8 13.25Zm0-1.85q.375 0 .638-.263T8.9 10.5q0-.375-.263-.638T8 9.6q-.375 0-.638.263T7.1 10.5q0 .375.263.638T8 11.4Zm9 7.6V5h2v14h-2Zm4 0V5h2v14h-2ZM8 10.5ZM8 17Z"/>`),
		g.Group(children),
	)
}
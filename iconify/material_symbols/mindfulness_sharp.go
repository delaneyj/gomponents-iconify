package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MindfulnessSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 14h2V7h-2v7Zm3-1.5h2V8h-2v4.5ZM8 12h2V8H8v4ZM6 22v-4.3q-1.425-1.3-2.212-3.038T3 11q0-3.75 2.625-6.375T12 2q3.125 0 5.538 1.838t3.137 4.787L22.3 15H19v5h-4v2H6Z"/>`),
		g.Group(children),
	)
}
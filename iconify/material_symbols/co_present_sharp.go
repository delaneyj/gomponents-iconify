package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoPresentSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 21V5H3v8H1V3h22v18h-2ZM9 14q-1.65 0-2.825-1.175T5 10q0-1.65 1.175-2.825T9 6q1.65 0 2.825 1.175T13 10q0 1.65-1.175 2.825T9 14Zm-8 8v-2.8q0-.85.438-1.562T2.6 16.55q1.55-.775 3.15-1.163T9 15q1.65 0 3.25.388t3.15 1.162q.725.375 1.163 1.088T17 19.2V22H1Z"/>`),
		g.Group(children),
	)
}
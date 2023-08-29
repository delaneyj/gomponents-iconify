package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Puzzle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 3c-2.21 0-4 1.79-4 4H7v5c-2.21 0-4 1.79-4 4s1.79 4 4 4v5h6c0 2.21 1.79 4 4 4s4-1.79 4-4h6v-7h-2c-1.19 0-2-.81-2-2s.81-2 2-2h2V7h-6c0-2.21-1.79-4-4-4zm0 2c1.19 0 2 .81 2 2v2h6v3c-2.21 0-4 1.79-4 4s1.79 4 4 4v3h-6v2c0 1.19-.81 2-2 2s-2-.81-2-2v-2H9v-5H7c-1.19 0-2-.81-2-2s.81-2 2-2h2V9h6V7c0-1.19.81-2 2-2z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TetrisAppFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 6a2 2 0 0 1 2-2h8v10H16V6ZM6 16a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h8V16H6Zm20 0H16v10h8a2 2 0 0 0 2-2v-8ZM36 4h-8v10h8a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2ZM10 35a2 2 0 0 1 2-2h8v10h-8a2 2 0 0 1-2-2v-6Zm22-2H22v10h10V33Zm2 0h10v8a2 2 0 0 1-2 2h-8V33Zm2-12a2 2 0 0 0-2 2v8h10v-8a2 2 0 0 0-2-2h-6Z"/>`),
		g.Group(children),
	)
}
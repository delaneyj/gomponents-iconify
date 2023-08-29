package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpenTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.5 4c.98 0 1.865.402 2.5 1.05A3.49 3.49 0 0 1 16.5 4H24a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2h-7.5a3.49 3.49 0 0 1-2.5-1.05A3.49 3.49 0 0 1 11.5 24H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h7.5ZM13 20.5v-13A1.5 1.5 0 0 0 11.5 6H4v16h7.5a1.5 1.5 0 0 0 1.5-1.5Zm2-13v13a1.5 1.5 0 0 0 1.5 1.5H24V6h-7.5A1.5 1.5 0 0 0 15 7.5Z"/>`),
		g.Group(children),
	)
}
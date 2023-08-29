package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoopLeftLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4a7.992 7.992 0 0 0-6.616 3.5H8v2H2v-6h2V6a9.984 9.984 0 0 1 8-4c5.523 0 10 4.477 10 10h-2a8 8 0 0 0-8-8Zm-8 8a8 8 0 0 0 14.616 4.5H16v-2h6v6h-2V18a9.984 9.984 0 0 1-8 4C6.477 22 2 17.523 2 12h2Z"/>`),
		g.Group(children),
	)
}
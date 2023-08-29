package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlurOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.432 6.843L1.394 2.805L2.808 1.39l19.799 19.8l-1.415 1.414l-3.038-3.04A9 9 0 0 1 5.432 6.845Zm2.811-2.817L12 .269l6.364 6.364a9.002 9.002 0 0 1 2.05 9.564L8.244 4.026Z"/>`),
		g.Group(children),
	)
}
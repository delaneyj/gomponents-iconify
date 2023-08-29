package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlurOffLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.154 19.565A9 9 0 0 1 5.432 6.843L1.394 2.805L2.808 1.39l19.799 19.8l-1.415 1.414l-3.038-3.04ZM6.848 8.259a7 7 0 0 0 9.89 9.89l-9.89-9.89Zm13.566 7.938l-1.598-1.599a6.996 6.996 0 0 0-1.866-6.55L12 3.097L9.658 5.44L8.243 4.026L12 .269l6.364 6.364a9.002 9.002 0 0 1 2.05 9.564Z"/>`),
		g.Group(children),
	)
}
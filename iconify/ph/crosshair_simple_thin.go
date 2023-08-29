package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrosshairSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm4 191.91V184a4 4 0 0 0-8 0v35.91A92.13 92.13 0 0 1 36.09 132H72a4 4 0 0 0 0-8H36.09A92.13 92.13 0 0 1 124 36.09V72a4 4 0 0 0 8 0V36.09A92.13 92.13 0 0 1 219.91 124H184a4 4 0 0 0 0 8h35.91A92.13 92.13 0 0 1 132 219.91Z"/>`),
		g.Group(children),
	)
}
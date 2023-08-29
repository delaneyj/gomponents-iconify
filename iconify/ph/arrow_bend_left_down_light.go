package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendLeftDownLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M206 32a6 6 0 0 1-6 6a90.1 90.1 0 0 0-90 90v81.51l37.76-37.75a6 6 0 0 1 8.48 8.48l-48 48a6 6 0 0 1-8.48 0l-48-48a6 6 0 0 1 8.48-8.48L98 209.51V128A102.12 102.12 0 0 1 200 26a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}
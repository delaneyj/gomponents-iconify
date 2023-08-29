package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLineDownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M72 40a8 8 0 0 1 8-8h96a8 8 0 0 1 0 16H80a8 8 0 0 1-8-8Zm159.39 92.94A8 8 0 0 0 224 128h-40V72a8 8 0 0 0-8-8H80a8 8 0 0 0-8 8v56H32a8 8 0 0 0-5.66 13.66l96 96a8 8 0 0 0 11.32 0l96-96a8 8 0 0 0 1.73-8.72Z"/>`),
		g.Group(children),
	)
}
package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.397 2.21A.726.726 0 0 1 8.41 2.2a.69.69 0 0 1 .012.99L2.373 9.243l16.931.105a.7.7 0 0 1-.008 1.4l-16.809-.104l6.303 6.16a.69.69 0 0 1 .076.903l-.075.087a.727.727 0 0 1-1.012 0L.21 10.4a.69.69 0 0 1-.005-.985Z"/>`),
		g.Group(children),
	)
}
package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.397 2.21A.726.726 0 0 1 9.41 2.2a.69.69 0 0 1 .012.99l-6.7 6.707l7.07 6.908a.69.69 0 0 1 0 .99a.727.727 0 0 1-1.012 0L1.21 10.4a.69.69 0 0 1-.005-.985Zm9 0a.726.726 0 0 1 1.012-.01a.69.69 0 0 1 .012.99l-6.7 6.707l7.07 6.908a.69.69 0 0 1 0 .99a.727.727 0 0 1-1.012 0L10.21 10.4a.69.69 0 0 1-.005-.985Z"/>`),
		g.Group(children),
	)
}
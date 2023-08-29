package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5.5v13m3.48-5.636l9.016 5.259A1 1 0 0 0 19 17.259V6.741a1 1 0 0 0-1.504-.864l-9.015 5.26a1 1 0 0 0 0 1.727Z"/>`),
		g.Group(children),
	)
}
package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoldCyrlBe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 19V1h12v3H8v4h1c2 0 3.62.14 4.72.61a5.27 5.27 0 0 1 2.48 1.95c.53.82.8 1.88.8 2.94c0 1.78-.78 3-2 4c-1.2.97-3.35 1.5-6 1.5H4Zm5-3a6.7 6.7 0 0 0 3.01-.68c.7-.37.99-.9.99-1.82c0-.96-.33-1.71-1.07-2.01A7.61 7.61 0 0 0 9 11H8v5h1Z"/>`),
		g.Group(children),
	)
}
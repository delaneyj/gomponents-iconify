package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.657 8.962l-1.418 1.411l-3.255-3.27l-.013 13.605l-2-.002l.013-13.568l-3.23 3.215l-1.41-1.417l5.67-5.644l5.643 5.67Z"/>`),
		g.Group(children),
	)
}
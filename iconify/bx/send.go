package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Send(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.426 11.095l-17-8A.999.999 0 0 0 3.03 4.242L4.969 12L3.03 19.758a.998.998 0 0 0 1.396 1.147l17-8a1 1 0 0 0 0-1.81zM5.481 18.197l.839-3.357L12 12L6.32 9.16l-.839-3.357L18.651 12l-13.17 6.197z"/>`),
		g.Group(children),
	)
}
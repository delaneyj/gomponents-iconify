package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zm3 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2zM5 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm3 9a4.999 4.999 0 0 1-4.288-2.427l1.286-.772C5.61 10.819 6.725 11.5 8 11.5s2.389-.681 3.002-1.699l1.286.772A4.996 4.996 0 0 1 8 13z"/>`),
		g.Group(children),
	)
}
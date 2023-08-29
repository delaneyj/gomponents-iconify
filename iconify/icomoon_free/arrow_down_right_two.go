package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRightTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2.293 3.707L10.586 12H7a1 1 0 0 0 0 2h6a1.002 1.002 0 0 0 1-1h.001V7a1 1 0 0 0-2 0v3.586L3.708 2.293a.997.997 0 0 0-1.414 0a.999.999 0 0 0 0 1.414z"/>`),
		g.Group(children),
	)
}
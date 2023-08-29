package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeftTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.293 2.293L4 10.586V7a1 1 0 0 0-2 0v6a1.002 1.002 0 0 0 1 1v.001h6a1 1 0 0 0 0-2H5.414l8.293-8.293a.997.997 0 0 0 0-1.414a.999.999 0 0 0-1.414 0z"/>`),
		g.Group(children),
	)
}
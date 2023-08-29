package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRightTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3.707 13.707L12 5.414V9a1 1 0 0 0 2 0V3a1.002 1.002 0 0 0-1-1v-.001H7a1 1 0 0 0 0 2h3.586l-8.293 8.293a.997.997 0 0 0 0 1.414a.999.999 0 0 0 1.414 0z"/>`),
		g.Group(children),
	)
}
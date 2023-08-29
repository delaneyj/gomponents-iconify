package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m13.707 6.293l-5-5a.999.999 0 0 0-1.414 0l-5 5a.999.999 0 1 0 1.414 1.414L7 4.414V14a1 1 0 0 0 2 0V4.414l3.293 3.293a.997.997 0 0 0 1.414 0a.999.999 0 0 0 0-1.414z"/>`),
		g.Group(children),
	)
}
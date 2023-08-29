package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14 4.213L7.5.42L1 4.213v6.574l6.5 3.792l6.5-3.792V4.213Z"/>`),
		g.Group(children),
	)
}
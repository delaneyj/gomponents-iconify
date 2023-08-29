package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.24 12.25l7.26 7.25v1.42l-7.96-7.96A5.247 5.247 0 0 1 2 9.25A5.25 5.25 0 0 1 7.25 4c1.75 0 3.3.85 4.25 2.17v.97h-.56A4.241 4.241 0 0 0 7.25 5A4.25 4.25 0 0 0 3 9.25c0 1.17.47 2.25 1.24 3Z"/>`),
		g.Group(children),
	)
}
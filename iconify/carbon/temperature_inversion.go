package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemperatureInversion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 8h-6v6h2v-2.459A5.939 5.939 0 0 1 28 16a6.004 6.004 0 0 1-5.095 5.924A7.027 7.027 0 0 0 20 17.262V6a4 4 0 0 0-8 0v11.262a6.996 6.996 0 1 0 10.93 6.68A7.98 7.98 0 0 0 27.264 10H30zM14 6a2 2 0 0 1 4 0v7h-4zm2 22a4.997 4.997 0 0 1-2.499-9.325l.499-.29V15h4v3.386l.499.289A4.997 4.997 0 0 1 16 28zM2 10a8.04 8.04 0 0 0 2.735 6H2v2h6v-6H6v2.46A5.94 5.94 0 0 1 4 10a6.007 6.007 0 0 1 6-6V2a8.01 8.01 0 0 0-8 8z"/>`),
		g.Group(children),
	)
}
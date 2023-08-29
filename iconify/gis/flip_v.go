package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M0 52.5h10v-5H0Zm15 0h10v-5H15Zm15 0h10v-5H30Zm15 0h10v-5H45Zm15 0h10v-5H60Zm15 0h10v-5H75Zm15 0h10v-5H90ZM26.055.002a3.5 3.5 0 0 0-3.098 2.05l-14.072 31A3.5 3.5 0 0 0 12.075 38l75.857-.05a3.5 3.5 0 0 0 2.812-5.58L70.131 4.474a3.5 3.5 0 0 0-2.555-1.41L26.404.01a3.5 3.5 0 0 0-.35-.008Zm2.267 7.17l37.135 2.754l15.54 21.027l-63.491.043ZM12.074 62a3.5 3.5 0 0 0-3.19 4.947l14.073 31a3.5 3.5 0 0 0 3.098 2.051a3.5 3.5 0 0 0 .35-.008l41.171-3.054a3.5 3.5 0 0 0 2.555-1.41L90.744 67.63a3.5 3.5 0 0 0-2.812-5.58L12.074 62z" color="currentColor"/>`),
		g.Group(children),
	)
}
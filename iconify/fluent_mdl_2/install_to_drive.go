package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstallToDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1728 640q66 0 124 25t101 69t69 102t26 124v448h-128V960q0-40-15-75t-41-61t-61-41t-75-15H320q-40 0-75 15t-61 41t-41 61t-15 75v320h1408v128H0V960q0-66 25-124t68-101t102-69t125-26h1408zm227 899l90 90l-317 317l-317-317l90-90l163 162v-549h128v549l163-162zm-163-515h-128V896h128v128zm-384-128h128v128h-128V896z"/>`),
		g.Group(children),
	)
}
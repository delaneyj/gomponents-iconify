package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloneToDesktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 1536h256v128H640v-128h256v-128H0V256h1920v768h-128V384H128v896h1408v128h-512v128zm1021 93l-317 317l-317-317l90-90l163 162v-549h128v549l163-162l90 90z"/>`),
		g.Group(children),
	)
}
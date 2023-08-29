package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoTwoAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1408 576q0-26 19-45t45-19q26 0 45 19t19 45q0 26-19 45t-45 19q-26 0-45-19t-19-45zm640 960v128h-256v256h-128v-256h-256v-128h256v-256h128v256h256zM922 1216l358 357v91H0V256h1792v896h-128V384H128v421l192-191l512 512l256-256l448 447v91h-91l-357-358l-166 166zM320 794L128 987v549h933L320 794z"/>`),
		g.Group(children),
	)
}
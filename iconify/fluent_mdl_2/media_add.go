package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 384h128v128H256V384zm0 256h128v128H256V640zm896-256h128v128h-128V384zm731 896h-310l155-154l155 154zm-475-320q0-26 19-45t45-19q26 0 45 19t19 45q0 26-19 45t-45 19q-26 0-45-19t-19-45zm640-320v896h-128V768H640v549l320-319l448 447l128-128v182l-37 37h-182l-357-358l-320 321v165h640v128H512v-640H0V0h1536v640h512zm-640 0V128h-128v128h-128V128H384v128H256V128H128v896h128V896h128v128h128V640h896zm384 1024h256v128h-256v256h-128v-256h-256v-128h256v-256h128v256z"/>`),
		g.Group(children),
	)
}
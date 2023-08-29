package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FixedColumnWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 0h128v768H0V0h128v640h256V384h128v256h256V128h128v512h256V384h128v256h256V128h128v512h256V0zM0 2048V896h2048v1152h-128v-128h-512v128h-128v-128H768v128H640v-128H128v128H0zm768-256h512v-384H768v384zm640 0h512v-384h-512v384zM128 1024v256h1792v-256H128zm512 384H128v384h512v-384z"/>`),
		g.Group(children),
	)
}
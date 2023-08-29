package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1408 1600v-320h128v320h-128zm0 192v-128h128v128h-128zM0 128h2048v1434l-128-256V256H128v677l448-447l572 572l-60 120l-512-512l-448 449v421h781l-64 128H0V128zm1600 384q-26 0-45-19t-19-45q0-26 19-45t45-19q26 0 45 19t19 45q0 26-19 45t-45 19zM896 1920l576-1152l576 1152H896zm181-112h790l-395-790l-395 790z"/>`),
		g.Group(children),
	)
}
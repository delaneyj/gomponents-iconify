package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartVerticalEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1408 384v641l-128 128V256h384v531q-39 22-67 50t-61 61V384h-128zM640 512v1152H256V512h384zM512 1536V640H384v896h128zm1536-442q0 39-15 76t-43 65l-717 717l-377 94l94-377l717-716q29-29 64-43t77-14q42 0 78 15t64 41t42 63t16 79zm-128 0q0-32-20-51t-52-19q-14 0-27 4t-23 15l-692 692l-34 135l135-34l692-691q21-21 21-51zM896 1536h1l-25 25q-12 12-27 26q-5 20-9 39t-10 38h-58V768h384v513l-127 127h-1V896H896v640zM128 128v1664h666l-32 128H0V128h128z"/>`),
		g.Group(children),
	)
}
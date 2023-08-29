package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dataflows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1280 0h640v640h-256v512H512v256h256v640H128v-640h256V640H128V0h640v640H512v384h1024V640h-256V0zM256 512h384V128H256v384zm384 1408v-384H256v384h384zM1792 512V128h-384v384h384z"/>`),
		g.Group(children),
	)
}
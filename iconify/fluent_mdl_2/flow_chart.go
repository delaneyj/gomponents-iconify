package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1408 1152h640v640h-640v-256H992l-416 416l-480-480l416-416V640H256V0h640v640H640v416l352 352h416v-256zM384 128v384h384V128H384zm192 1632l288-288l-288-288l-288 288l288 288zm1344-96v-384h-384v384h384z"/>`),
		g.Group(children),
	)
}
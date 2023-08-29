package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexaditeInvestigationCancel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m768 1169l197 113q9 7 28 25t40 40t40 41t29 30l-78 45l-384-220V805l384-220l384 220v286l-128-128v-84l-256-146l-256 146v290zM256 586v876l742 424l-94 93l-776-443V512L1024 0l896 512v451l-128 128V586l-768-439l-768 439zm1773 675l-339 339l339 339l-90 90l-339-339l-339 339l-90-90l339-339l-339-339l90-90l339 339l339-339l90 90z"/>`),
		g.Group(children),
	)
}
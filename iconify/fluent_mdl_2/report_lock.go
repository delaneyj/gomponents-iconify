package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReportLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1408h128v640h-768v-640h128v-128q0-53 20-99t55-82t81-55t100-20q53 0 99 20t82 55t55 81t20 100v128zm-384-128v128h256v-128q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50zm384 640v-384h-512v384h512zM896 1792H640v-384h256v384zm-384 0H256v-640h256v640zM0 0h987l549 549v347h-128V640H896V128H128v1792h1024v128H0V0zm1024 219v293h293l-293-293zm256 677v384h-128v512h-128V896h256z"/>`),
		g.Group(children),
	)
}
package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReportDocument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 549v1499H256V0h987l549 549zm-512-37h293l-293-293v293zm384 1408V640h-512V128H384v1792h1280zm-768-512h256v384H896v-384zm-384-256h256v640H512v-640zm768-256h256v896h-256V896z"/>`),
		g.Group(children),
	)
}
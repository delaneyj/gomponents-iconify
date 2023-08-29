package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadDocument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1920h1024v128H128V0h1115l549 549v347h-128V640h-512V128H256v1792zM1280 512h293l-293-293v293zm512 1061l163-162l90 90l-317 317l-317-317l90-90l163 162v-549h128v549zm256 347v128h-640v-128h640z"/>`),
		g.Group(children),
	)
}
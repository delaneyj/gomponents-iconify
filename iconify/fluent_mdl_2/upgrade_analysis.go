package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpgradeAnalysis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1277 1251l-90 90l-163-163v550H896v-550l-163 163l-90-90l317-318l317 318zm752 202l-557 558l-269-270l90-90l179 178l467-466l90 90zM256 1920h896v128H128V0h1115l549 549v734l-128 128V640h-512V128H256v1792zM1280 512h293l-293-293v293z"/>`),
		g.Group(children),
	)
}
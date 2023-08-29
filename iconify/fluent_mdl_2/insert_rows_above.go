package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertRowsAbove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1664H0V128h512L384 256H128v384h640v512h384V640h768V256h-384l-128-128h640zM640 1280H128v384h512v-384zm640 0H768v384h512v-384zm640 0h-512v384h512v-384zM621 525l-90-90L960 6l429 429l-90 90l-275-275v774H896V250L621 525z"/>`),
		g.Group(children),
	)
}
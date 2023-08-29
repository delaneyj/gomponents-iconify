package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimelineProgress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 640h640v128H256V640zm1024 512h512v128h-512v-128zm768-896v1408H0V256h2048zm-127 128h-514v384h-127V384H768v128H640V384H128v1153h512V896h128v641h512v-129h127v129h514V384zM897 896h639v128H897V896z"/>`),
		g.Group(children),
	)
}
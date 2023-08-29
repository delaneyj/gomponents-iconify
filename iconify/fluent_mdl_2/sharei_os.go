package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareiOS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 250L749 525l-90-90L1088 6l429 429l-90 90l-275-275v1158h-128V250zm256 390h512v1408H384V640h512v128H512v1152h1152V768h-384V640z"/>`),
		g.Group(children),
	)
}
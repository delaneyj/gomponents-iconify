package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 0h2048L1024 2048L0 0zm1841 128H207l128 256h1378l128-256zM783 1280l241 482l241-482H783zm546-128l128-256H591l128 256h610zm192-384l128-256H399l128 256h994z"/>`),
		g.Group(children),
	)
}
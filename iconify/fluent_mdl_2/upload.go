package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 0v128H384V0h1152zm45 979l-90 90l-467-470v1449H896V599l-467 470l-90-90l621-626l621 626z"/>`),
		g.Group(children),
	)
}
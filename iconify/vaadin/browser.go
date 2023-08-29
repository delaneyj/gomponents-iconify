package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Browser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 1V0H0v15h1v1h15V1h-1zM3 1h9v1H3V1zM1 1h1v1H1V1zm0 2h13v11H1V3z"/>`),
		g.Group(children),
	)
}
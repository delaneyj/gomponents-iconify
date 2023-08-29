package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 0H0v16h16V0zM1 5V1h14v4H1zm0 5V6h14v4H1zm0 5v-4h14v4H1z"/>`),
		g.Group(children),
	)
}
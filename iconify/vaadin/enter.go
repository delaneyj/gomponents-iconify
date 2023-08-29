package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Enter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 0v6H1v10h14V0H4zm8 11H7v2l-3-2.5L7 8v2h4V7h1v4z"/>`),
		g.Group(children),
	)
}
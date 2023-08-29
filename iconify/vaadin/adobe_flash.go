package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeFlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0v16h16V0H0zm13 4.4C10 4.4 9.7 7 9.7 7H11v2H8.6C6.8 14.8 3 14 3 14v-2.5s2.5.6 3.9-4C8.7 1.4 13 2 13 2v2.4z"/>`),
		g.Group(children),
	)
}
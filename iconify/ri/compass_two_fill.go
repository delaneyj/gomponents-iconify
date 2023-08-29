package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.328 4.258L10.586 12L12 13.414l7.742-7.742A9.957 9.957 0 0 1 22 12c0 5.52-4.48 10-10 10S2 17.52 2 12S6.48 2 12 2c2.4 0 4.604.847 6.328 2.258Z"/>`),
		g.Group(children),
	)
}
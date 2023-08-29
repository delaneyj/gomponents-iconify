package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOuterLightDashedUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 0h4v2H1V0m6 0h3v2H7V0m5 0h4v2h-4V0m6 0h3v2h-3V0Z"/>`),
		g.Group(children),
	)
}
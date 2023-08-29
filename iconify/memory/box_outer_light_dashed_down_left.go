package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOuterLightDashedDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 18h2v2h2v2H0v-4m6 2h4v2H6v-2m6 0h4v2h-4v-2m6 0h3v2h-3v-2M0 16v-4h2v4H0m0-6V7h2v3H0m0-5V1h2v4H0Z"/>`),
		g.Group(children),
	)
}
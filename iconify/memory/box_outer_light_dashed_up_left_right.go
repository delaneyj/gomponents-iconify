package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOuterLightDashedUpLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 4h-2V2h-2V0h4v4m-6-2h-4V0h4v2m-6 0H6V0h4v2M4 2H2v2H0V0h4v2m18 4v4h-2V6h2m0 6v3h-2v-3h2m0 5v4h-2v-4h2M0 21v-4h2v4H0m0-6v-3h2v3H0m0-5V6h2v4H0Z"/>`),
		g.Group(children),
	)
}
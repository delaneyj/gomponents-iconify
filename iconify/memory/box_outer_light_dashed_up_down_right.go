package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOuterLightDashedUpDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 22v-2h2v-2h2v4h-4m2-6v-4h2v4h-2m0-6V6h2v4h-2m0-6V2h-2V0h4v4h-2m-4 18h-4v-2h4v2m-6 0H7v-2h3v2m-5 0H1v-2h4v2M1 0h4v2H1V0m6 0h3v2H7V0m5 0h4v2h-4V0Z"/>`),
		g.Group(children),
	)
}
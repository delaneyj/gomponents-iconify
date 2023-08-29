package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOuterLightDashedUpDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M21 2h-4V0h4v2m-6 0h-3V0h3v2m-5 0H6V0h4v2M4 2H1V0h3v2m17 20h-4v-2h4v2m-6 0h-3v-2h3v2m-5 0H6v-2h4v2m-6 0H1v-2h3v2Z"/>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aquarius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 20L14 8l6 9l8-9l7 11l9-9.957M4 40l10-12l6 9l8-9l7 11l9-9.956"/>`),
		g.Group(children),
	)
}
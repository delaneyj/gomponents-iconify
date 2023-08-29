package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSpaceDashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21H3V3h8v18zm2 0h8v-9h-8v9zm8-11V3h-8v7h8z"/>`),
		g.Group(children),
	)
}
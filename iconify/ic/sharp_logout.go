package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpLogout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5h7V3H3v18h9v-2H5z"/><path fill="currentColor" d="m21 12l-4-4v3H9v2h8v3z"/>`),
		g.Group(children),
	)
}
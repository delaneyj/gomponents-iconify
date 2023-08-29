package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchLayerTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 26v-2H5.83l2.58-2.59L7 20l-5 5l5 5l1.41-1.41L5.83 26H16zm0-12v-2H5.83l2.58-2.59L7 8l-5 5l5 5l1.41-1.41L5.83 14H16zm0-8v2h10.17l-2.58 2.59L25 12l5-5l-5-5l-1.41 1.41L26.17 6H16zm0 12v2h10.17l-2.58 2.59L25 24l5-5l-5-5l-1.41 1.41L26.17 18H16z"/>`),
		g.Group(children),
	)
}
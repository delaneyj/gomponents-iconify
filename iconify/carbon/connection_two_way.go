package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectionTwoWay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 8h2v4H8zm0 6h2v4H8zm6-6h2v4h-2zm0 6h2v4h-2zm-6 6h2v4H8zm6 0h2v4h-2zm16 2h-8.17l2.58 2.59L23 26l-5-5l5-5l1.41 1.41L21.83 20H30v2zM19 12h8.17l-2.58 2.59L26 16l5-5l-5-5l-1.41 1.41L27.17 10H19v2z"/><path fill="currentColor" d="M28 26v2H4V4h16v2h2V4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v26h28v-4Z"/>`),
		g.Group(children),
	)
}
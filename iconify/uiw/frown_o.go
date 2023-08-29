package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrownO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0Zm0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21ZM10 11.4a3.637 3.637 0 0 1 3.607 3.175a.682.682 0 1 1-1.352.172a2.273 2.273 0 0 0-4.511.015a.682.682 0 1 1-1.354-.163a3.637 3.637 0 0 1 3.61-3.2ZM5.814 6.28a1.395 1.395 0 1 1 0 2.79a1.395 1.395 0 0 1 0-2.79Zm8.372 0a1.395 1.395 0 1 1 0 2.79a1.395 1.395 0 0 1 0-2.79Z"/>`),
		g.Group(children),
	)
}
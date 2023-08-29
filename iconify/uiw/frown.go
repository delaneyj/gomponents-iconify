package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Frown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0Zm0 11.399a3.637 3.637 0 0 0-3.61 3.2a.682.682 0 0 0 1.354.162a2.273 2.273 0 0 1 4.51-.015a.682.682 0 1 0 1.353-.172A3.637 3.637 0 0 0 10 11.4Zm-4.186-5.12a1.395 1.395 0 1 0 0 2.79a1.395 1.395 0 0 0 0-2.79Zm8.372 0a1.395 1.395 0 1 0 0 2.79a1.395 1.395 0 0 0 0-2.79Z"/>`),
		g.Group(children),
	)
}
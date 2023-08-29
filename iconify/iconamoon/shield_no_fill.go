package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldNoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.606 2.08a1 1 0 0 1 .788 0l6.394 2.741A2 2 0 0 1 20 6.66v6.86a7 7 0 0 1-3.527 6.077l-3.977 2.272a1 1 0 0 1-.992 0l-3.977-2.272A7 7 0 0 1 4 13.518V6.66a2 2 0 0 1 1.212-1.838l6.394-2.74Zm-.899 7.213a1 1 0 0 0-1.414 1.414L10.586 12l-1.293 1.293a1 1 0 1 0 1.414 1.414L12 13.414l1.293 1.293a1 1 0 0 0 1.414-1.414L13.414 12l1.293-1.293a1 1 0 0 0-1.414-1.414L12 10.586l-1.293-1.293Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
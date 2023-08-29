package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseCircleOneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm7.293-2.707a1 1 0 0 1 1.414 0L12 10.586l1.293-1.293a1 1 0 1 1 1.414 1.414L13.414 12l1.293 1.293a1 1 0 0 1-1.414 1.414L12 13.414l-1.293 1.293a1 1 0 0 1-1.414-1.414L10.586 12l-1.293-1.293a1 1 0 0 1 0-1.414Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
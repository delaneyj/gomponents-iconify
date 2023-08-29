package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopRightOneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18 17V7a1 1 0 0 0-1-1H7a1 1 0 0 0-.707 1.707L10.586 12l-4.293 4.293a1 1 0 1 0 1.414 1.414L12 13.414l4.293 4.293A1 1 0 0 0 18 17Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
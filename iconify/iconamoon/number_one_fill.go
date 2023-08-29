package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberOneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.383 3.076A1 1 0 0 1 13 4v16a1 1 0 1 1-2 0V6.414L9.707 7.707a1 1 0 0 1-1.414-1.414l3-3a1 1 0 0 1 1.09-.217Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
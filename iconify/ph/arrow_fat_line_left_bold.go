package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLineLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 68h-44V32a12 12 0 0 0-20.49-8.49l-96 96a12 12 0 0 0 0 17l96 96A12 12 0 0 0 132 224v-36h44a12 12 0 0 0 12-12V80a12 12 0 0 0-12-12Zm-12 96h-44a12 12 0 0 0-12 12v19l-67-67l67-67v19a12 12 0 0 0 12 12h44Zm64-84v96a12 12 0 0 1-24 0V80a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}
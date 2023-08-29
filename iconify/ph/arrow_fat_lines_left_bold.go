package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLinesLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M136 68h-4V32a12 12 0 0 0-20.49-8.49l-96 96a12 12 0 0 0 0 17l96 96A12 12 0 0 0 132 224v-36h4a12 12 0 0 0 12-12V80a12 12 0 0 0-12-12Zm-12 96h-4a12 12 0 0 0-12 12v19l-67-67l67-67v19a12 12 0 0 0 12 12h4Zm104-84v96a12 12 0 0 1-24 0V80a12 12 0 0 1 24 0Zm-40 0v96a12 12 0 0 1-24 0V80a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLinesLeftLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M152 74h-26V32a6 6 0 0 0-10.24-4.24l-96 96a6 6 0 0 0 0 8.48l96 96A6 6 0 0 0 126 224v-42h26a6 6 0 0 0 6-6V80a6 6 0 0 0-6-6Zm-6 96h-26a6 6 0 0 0-6 6v33.51L32.49 128L114 46.49V80a6 6 0 0 0 6 6h26Zm76-90v96a6 6 0 0 1-12 0V80a6 6 0 0 1 12 0Zm-32 0v96a6 6 0 0 1-12 0V80a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLineLeftLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 74h-58V32a6 6 0 0 0-10.24-4.24l-96 96a6 6 0 0 0 0 8.48l96 96A6 6 0 0 0 126 224v-42h58a6 6 0 0 0 6-6V80a6 6 0 0 0-6-6Zm-6 96h-58a6 6 0 0 0-6 6v33.51L32.49 128L114 46.49V80a6 6 0 0 0 6 6h58Zm44-90v96a6 6 0 0 1-12 0V80a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellSignalFullLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M166 72v128a6 6 0 0 1-12 0V72a6 6 0 0 1 12 0Zm34-46a6 6 0 0 0-6 6v168a6 6 0 0 0 12 0V32a6 6 0 0 0-6-6Zm-80 80a6 6 0 0 0-6 6v88a6 6 0 0 0 12 0v-88a6 6 0 0 0-6-6Zm-40 40a6 6 0 0 0-6 6v48a6 6 0 0 0 12 0v-48a6 6 0 0 0-6-6Zm-40 40a6 6 0 0 0-6 6v8a6 6 0 0 0 12 0v-8a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}
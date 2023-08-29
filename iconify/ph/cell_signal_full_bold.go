package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellSignalFullBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M172 72v128a12 12 0 0 1-24 0V72a12 12 0 0 1 24 0Zm28-52a12 12 0 0 0-12 12v168a12 12 0 0 0 24 0V32a12 12 0 0 0-12-12Zm-80 80a12 12 0 0 0-12 12v88a12 12 0 0 0 24 0v-88a12 12 0 0 0-12-12Zm-40 40a12 12 0 0 0-12 12v48a12 12 0 0 0 24 0v-48a12 12 0 0 0-12-12Zm-40 40a12 12 0 0 0-12 12v8a12 12 0 0 0 24 0v-8a12 12 0 0 0-12-12Z"/>`),
		g.Group(children),
	)
}
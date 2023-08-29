package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m213.66 82.34l-56-56A8 8 0 0 0 152 24H56a16 16 0 0 0-16 16v176a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16V88a8 8 0 0 0-2.34-5.66ZM200 216H56V40h92.69L200 91.31V216Zm-16-96v72a8 8 0 0 1-8 8h-20a4 4 0 0 1-4-4v-44a8 8 0 0 0-8.53-8a8.17 8.17 0 0 0-7.47 8.25V196a4 4 0 0 1-4 4h-8a4 4 0 0 1-4-4v-44a8 8 0 0 0-8.53-8a8.17 8.17 0 0 0-7.47 8.25V196a4 4 0 0 1-4 4H80a8 8 0 0 1-8-8v-72a8 8 0 0 1 8-8h96a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}
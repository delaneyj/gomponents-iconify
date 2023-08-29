package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExportFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M136 96h-16V43.31L93.66 69.66a8 8 0 0 1-11.32-11.32l40-40a8 8 0 0 1 11.32 0l40 40a8 8 0 0 1-11.32 11.32L136 43.31Zm64 0h-64v40a8 8 0 0 1-16 0V96H56a16 16 0 0 0-16 16v96a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16v-96a16 16 0 0 0-16-16Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoAwesomeMosaicOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21H5q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h6v18Zm-2-2V5H5v14h4Zm4-8V3h6q.825 0 1.413.588T21 5v6h-8Zm2-2h4V5h-4v4Zm-2 12v-8h8v6q0 .825-.588 1.413T19 21h-6Zm2-2h4v-4h-4v4Zm-6-7Zm6-3Zm0 6Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PermMedia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.825 0-1.413-.588T1 19V6h2v13h17v2H3Zm4-4q-.825 0-1.413-.588T5 15V4q0-.825.588-1.413T7 2h5l2 2h7q.825 0 1.413.588T23 6v9q0 .825-.588 1.413T21 17H7Zm2-4h10l-3.45-4.5l-2.3 3l-1.55-2L9 13Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SixtyFpsSelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 14H6q-.825 0-1.413-.588T4 12V6q0-.825.588-1.413T6 4h5v2H6v2h3q.825 0 1.413.588T11 10v2q0 .825-.588 1.413T9 14Zm-3-4v2h3v-2H6Zm12 4h-3q-.825 0-1.413-.588T13 12V6q0-.825.588-1.413T15 4h3q.825 0 1.413.588T20 6v6q0 .825-.588 1.413T18 14Zm0-2V6h-3v6h3ZM3 22v-5h2v5H3Zm4 0v-5h2v5H7Zm4 0v-5h2v5h-2Zm4 0v-5h6v5h-6Z"/>`),
		g.Group(children),
	)
}
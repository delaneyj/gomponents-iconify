package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThirtyFpsSelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 14v-2h5v-2H4V8h5V6H4V4h5q.825 0 1.413.588T11 6v1.5q0 .65-.425 1.075T9.5 9q.65 0 1.075.425T11 10.5V12q0 .825-.588 1.413T9 14H4Zm14 0h-3q-.825 0-1.413-.588T13 12V6q0-.825.588-1.413T15 4h3q.825 0 1.413.588T20 6v6q0 .825-.588 1.413T18 14Zm0-2V6h-3v6h3ZM3 22v-5h2v5H3Zm4 0v-5h2v5H7Zm4 0v-5h2v5h-2Zm4 0v-5h6v5h-6Z"/>`),
		g.Group(children),
	)
}
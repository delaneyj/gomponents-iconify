package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.825 0-1.413-.588T1 19V6q0-.825.588-1.413T3 4h18q.825 0 1.413.588T23 6v13q0 .825-.588 1.413T21 21H3Zm0-2h18V6H3v13Zm5-2h8v-1H8v1Zm-3-3h2v-2H5v2Zm4 0h2v-2H9v2Zm4 0h2v-2h-2v2Zm4 0h2v-2h-2v2ZM5 10h2V8H5v2Zm4 0h2V8H9v2Zm4 0h2V8h-2v2Zm4 0h2V8h-2v2ZM3 19V6v13Z"/>`),
		g.Group(children),
	)
}
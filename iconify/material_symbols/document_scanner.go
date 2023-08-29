package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentScanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 6V1h5v2H4v3H2Zm18 0V3h-3V1h5v5h-2ZM2 23v-5h2v3h3v2H2Zm15 0v-2h3v-3h2v5h-5ZM7 20q-.825 0-1.413-.588T5 18V6q0-.825.588-1.413T7 4h10q.825 0 1.413.588T19 6v12q0 .825-.588 1.413T17 20H7Zm2-10h6V8H9v2Zm0 3h6v-2H9v2Zm0 3h6v-2H9v2Z"/>`),
		g.Group(children),
	)
}
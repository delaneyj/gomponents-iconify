package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20h2v2Zm-2-3.5v-2h2v2H3ZM3 15v-2h2v2H3Zm0-3.5v-2h2v2H3ZM3 8q0-.825.588-1.413T5 6v2H3Zm3.5 14v-2h2v2h-2ZM9 18q-.825 0-1.413-.588T7 16V4q0-.825.588-1.413T9 2h9q.825 0 1.413.588T20 4v12q0 .825-.588 1.413T18 18H9Zm0-2h9V4H9v12Zm1 6v-2h2v2h-2Zm3.5 0v-2h2q0 .825-.588 1.413T13.5 22Z"/>`),
		g.Group(children),
	)
}
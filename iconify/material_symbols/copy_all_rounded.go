package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyAllRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 18q-.825 0-1.413-.588T7 16V4q0-.825.588-1.413T9 2h9q.825 0 1.413.588T20 4v12q0 .825-.588 1.413T18 18H9Zm0-2h9V4H9v12Zm-6-1h2v-2H3v2Zm0-3.5h2v-2H3v2ZM10 22h2v-2h-2v2Zm-7-3.5h2v-2H3v2ZM5 22v-2H3q0 .825.588 1.413T5 22Zm1.5 0h2v-2h-2v2Zm7 0q.825 0 1.413-.588T15.5 20h-2v2ZM3 8h2V6q-.825 0-1.413.588T3 8Z"/>`),
		g.Group(children),
	)
}
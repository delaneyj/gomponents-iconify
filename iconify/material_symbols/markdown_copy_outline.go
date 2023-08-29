package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkdownCopyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 18q-.825 0-1.413-.588T7 16V4q0-.825.588-1.413T9 2h9q.825 0 1.413.588T20 4v12q0 .825-.588 1.413T18 18H9Zm0-2h9V4H9v12Zm-4 6q-.825 0-1.413-.588T3 20V6h2v14h11v2H5Zm5.25-9h1.5V8.5h1v3h1.5v-3h1V13h1.5V8q0-.425-.288-.713T15.75 7h-4.5q-.425 0-.713.288T10.25 8v5ZM9 16V4v12Z"/>`),
		g.Group(children),
	)
}
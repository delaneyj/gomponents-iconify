package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareWindows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 15V9q0-.825.588-1.413T9 7h8.15l-2.575-2.575L16 3l5 5l-5 5l-1.425-1.4L17.15 9H9v6H7Zm-2 6q-.825 0-1.413-.588T3 19V4h2v15h12v-4h2v4q0 .825-.588 1.413T17 21H5Z"/>`),
		g.Group(children),
	)
}
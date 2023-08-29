package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Road(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m5.469 6l-.188.75l-4.656 18L.281 26H31.72l-.344-1.25l-4.656-18l-.19-.75zM7.03 8h8.094l-.063 3h1.876l-.063-3h8.094l4.156 16H17.281l-.093-4h-2.375l-.094 4H2.875zM15 13l-.125 5h2.25L17 13z"/>`),
		g.Group(children),
	)
}
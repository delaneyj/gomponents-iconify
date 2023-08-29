package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-3h2V6q0-.825.588-1.413T6 4h15v2H6v11h6v3H2Zm13 0q-.425 0-.713-.288T14 19V9q0-.425.288-.713T15 8h6q.425 0 .713.288T22 9v10q0 .425-.288.713T21 20h-6Zm1-3h4v-7h-4v7Zm0 0h4h-4Z"/>`),
		g.Group(children),
	)
}
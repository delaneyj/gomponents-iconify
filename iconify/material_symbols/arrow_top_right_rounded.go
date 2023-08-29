package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopRightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.175 10H7v9q0 .425-.288.713T6 20q-.425 0-.713-.288T5 19v-9q0-.825.588-1.413T7 8h9.175L13.3 5.125q-.3-.3-.313-.713t.288-.712q.3-.3.713-.3t.712.3l4.6 4.6q.15.15.213.325t.062.375q0 .2-.063.375T19.3 9.7l-4.625 4.625q-.3.3-.7.288t-.7-.313q-.275-.3-.288-.7t.288-.7l2.9-2.9Z"/>`),
		g.Group(children),
	)
}
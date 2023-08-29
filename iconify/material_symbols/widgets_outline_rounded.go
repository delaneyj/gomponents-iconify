package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WidgetsOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.95 12.3L11.7 8.05q-.15-.15-.212-.325t-.063-.375q0-.2.063-.375t.212-.325l4.25-4.25q.15-.15.325-.212t.375-.063q.2 0 .375.063t.325.212l4.25 4.25q.15.15.213.325t.062.375q0 .2-.063.375t-.212.325l-4.25 4.25q-.15.15-.325.213t-.375.062q-.2 0-.375-.062t-.325-.213ZM3 10V4q0-.425.288-.713T4 3h6q.425 0 .713.288T11 4v6q0 .425-.288.713T10 11H4q-.425 0-.713-.288T3 10Zm10 10v-6q0-.425.288-.713T14 13h6q.425 0 .713.288T21 14v6q0 .425-.288.713T20 21h-6q-.425 0-.713-.288T13 20ZM3 20v-6q0-.425.288-.713T4 13h6q.425 0 .713.288T11 14v6q0 .425-.288.713T10 21H4q-.425 0-.713-.288T3 20ZM5 9h4V5H5v4Zm11.675 1.2L19.5 7.375L16.675 4.55L13.85 7.375l2.825 2.825ZM15 19h4v-4h-4v4ZM5 19h4v-4H5v4ZM9 9Zm4.85-1.625ZM9 15Zm6 0Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatColorTextRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 24q-.425 0-.713-.288T2 23v-2q0-.425.288-.713T3 20h18q.425 0 .713.288T22 21v2q0 .425-.288.713T21 24H3Zm4.125-7q-.575 0-.913-.488t-.137-1.037l4.4-11.725q.125-.35.425-.55t.65-.2h.9q.375 0 .663.2t.412.55L17.95 15.5q.2.55-.137 1.025T16.9 17q-.35 0-.65-.2t-.425-.55l-.975-2.85H9.2l-1.025 2.875q-.125.35-.413.537T7.126 17ZM9.9 11.4h4.2l-2.05-5.8h-.1L9.9 11.4Z"/>`),
		g.Group(children),
	)
}
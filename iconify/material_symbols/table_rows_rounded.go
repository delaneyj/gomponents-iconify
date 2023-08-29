package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableRowsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.425 0-.713-.288T3 20v-2.65q0-.425.288-.713T4 16.35h16q.425 0 .713.288t.287.712V20q0 .425-.288.713T20 21H4Zm0-6.65q-.425 0-.713-.288T3 13.35v-2.725q0-.425.288-.713T4 9.625h16q.425 0 .713.288t.287.712v2.725q0 .425-.288.713T20 14.35H4Zm0-6.725q-.425 0-.713-.288T3 6.625V4q0-.425.288-.713T4 3h16q.425 0 .713.288T21 4v2.625q0 .425-.288.713T20 7.625H4Z"/>`),
		g.Group(children),
	)
}
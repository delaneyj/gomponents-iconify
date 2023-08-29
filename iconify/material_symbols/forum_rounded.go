package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForumRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18q-.425 0-.713-.288T6 17v-2h13V6h2q.425 0 .713.288T22 7v12.575q0 .675-.613.938T20.3 20.3L18 18H7Zm-1-5l-2.3 2.3q-.475.475-1.088.213T2 14.575V3q0-.425.288-.713T3 2h13q.425 0 .713.288T17 3v9q0 .425-.288.713T16 13H6Z"/>`),
		g.Group(children),
	)
}
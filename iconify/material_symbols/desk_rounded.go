package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 17V7q0-.425.288-.713T3 6h18q.425 0 .713.288T22 7v10q0 .425-.288.713T21 18q-.2 0-.388-.075t-.324-.213q-.138-.137-.213-.324T20 17v-1h-4v1q0 .2-.075.388t-.213.324q-.137.138-.324.213T15 18q-.425 0-.712-.288T14 17V8H4v9q0 .425-.288.713T3 18q-.425 0-.713-.288T2 17Zm14-7h4V8h-4v2Zm0 4h4v-2h-4v2Z"/>`),
		g.Group(children),
	)
}
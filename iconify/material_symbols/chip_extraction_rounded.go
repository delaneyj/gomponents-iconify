package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChipExtractionRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 12q0 2.65 1.725 4.625t4.35 2.325q.4.05.663.35T12 20q0 .425-.363.7t-.812.225q-3.375-.425-5.6-2.963T3 12q0-3.4 2.212-5.938T10.8 3.075q.475-.05.838.213T12 4q0 .4-.263.7t-.662.35q-2.625.35-4.35 2.325T5 12Zm12.175 1H10q-.425 0-.713-.288T9 12q0-.425.288-.713T10 11h7.175L15.3 9.125q-.3-.3-.3-.713t.3-.712q.3-.3.7-.3t.7.3l3.6 3.6q.3.3.3.7t-.3.7l-3.6 3.6q-.3.3-.7.288t-.7-.313q-.3-.3-.3-.7t.3-.7L17.175 13Z"/>`),
		g.Group(children),
	)
}
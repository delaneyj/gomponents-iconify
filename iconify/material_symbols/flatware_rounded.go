package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlatwareRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21q-.425 0-.713-.288T12 20v-9.1q-1.05-.5-1.525-1.563T10 7.1q0-1.575.788-2.837T13 3q1.425 0 2.212 1.263T16 7.1q0 1.175-.475 2.238T14 10.9V20q0 .425-.288.713T13 21Zm5 0q-.425 0-.713-.288T17 20V4.25q0-.55.375-.875t.85-.175q1.2.4 1.988 1.425T21 7v5q0 .425-.287.713T20 13h-1v7q0 .425-.288.713T18 21ZM6 21q-.425 0-.713-.288T5 20v-9q-.825 0-1.413-.588T3 9V3.7q0-.3.2-.5t.5-.2q.3 0 .5.2t.2.5V7h.9V3.7q0-.3.2-.5T6 3q.3 0 .5.2t.2.5V7h.9V3.7q0-.3.2-.5t.5-.2q.3 0 .5.2t.2.5V9q0 .825-.588 1.413T7 11v9q0 .425-.288.713T6 21Z"/>`),
		g.Group(children),
	)
}
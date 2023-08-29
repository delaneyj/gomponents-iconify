package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorSensorOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 15V5v14v-4Zm-2 1h2v3h6V5H9v9H7V3h10v18H7v-5Zm5-6q.425 0 .713-.288T13 9q0-.425-.288-.713T12 8q-.425 0-.713.288T11 9q0 .425.288.713T12 10Zm7-7h2v6h-2V3Zm-6 13H4v-2h9v2Z"/>`),
		g.Group(children),
	)
}
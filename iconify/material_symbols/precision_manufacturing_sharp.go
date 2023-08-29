package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrecisionManufacturingSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.975 21v-3h3.1l-2.55-8.35q-.675-.375-1.113-1.1T2.976 7q0-1.25.875-2.125T5.975 4q.975 0 1.738.563T8.774 6h3.2V4h1.8v.4L16 2.3l5.25 2.475l-.65 1.3L16.325 4.1l-2.35 2.2v1.4l2.35 2.15l4.3-1.925l.6 1.35L16 11.7l-2.225-2.1v.4h-1.8V8h-3.2q-.075.2-.163.375t-.237.375l5 9.25h3.6v3h-13Zm2-13q.425 0 .713-.288T6.974 7q0-.425-.287-.713T5.975 6q-.425 0-.712.288T4.974 7q0 .425.288.713T5.974 8Z"/>`),
		g.Group(children),
	)
}
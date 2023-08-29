package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrecisionManufacturing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.975 21v-3h3.1l-2.55-8.35q-.675-.375-1.113-1.1T2.976 7q0-1.25.875-2.125T5.975 4q.975 0 1.738.563T8.774 6h3.2V5q0-.425.288-.713T12.975 4q.225 0 .438.1t.362.3l1.7-1.6q.225-.225.538-.288t.612.088l3.9 1.8q.3.15.413.438t-.013.562q-.15.3-.438.388t-.562-.038l-3.6-1.65l-2.35 2.2v1.4l2.35 2.15l3.6-1.65q.275-.125.575-.025t.425.375q.15.3.025.575t-.425.425l-3.9 1.85q-.3.15-.613.087t-.537-.287l-1.7-1.6q-.15.15-.362.275t-.438.125q-.425 0-.712-.287T11.975 9V8h-3.2q-.075.2-.163.375t-.237.375l5 9.25h3.6v3h-13Zm2-13q.425 0 .713-.288T6.974 7q0-.425-.287-.713T5.975 6q-.425 0-.712.288T4.974 7q0 .425.288.713T5.974 8Z"/>`),
		g.Group(children),
	)
}
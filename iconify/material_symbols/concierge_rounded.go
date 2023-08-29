package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConciergeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22q-.425 0-.713-.288T10 21q0-.425.288-.713T11 20h11q.425 0 .713.288T23 21q0 .425-.288.713T22 22H11Zm0-3q0-2.025 1.275-3.538T15.5 13.6v-.625q0-.425.288-.712t.712-.288q.425 0 .713.288t.287.712v.625q1.925.35 3.213 1.863T22 19H11ZM1 7.5V11q0 .825.588 1.413T3 13q.825 0 1.413-.588T5 11V4q0-.825-.588-1.413T3 2q-.825 0-1.413.588T1 4v3.5Zm6 .975V11.4q0 .675.463 1.137T8.6 13h.025q.2 0 .363-.038t.337-.087l6.7-2.5q.425-.175.7-.563T17 8.95q0-.4-.275-.675T16.05 8H13l-1.3.5q-.2.075-.375 0t-.25-.275q-.075-.2.013-.388t.287-.262L13 7h7q.8 0 1.4-.575T22 5l-7.4-2.775q-.3-.125-.612-.138t-.613.088l-4.9 1.375q-.65.2-1.063.725T7 5.475v3Z"/>`),
		g.Group(children),
	)
}
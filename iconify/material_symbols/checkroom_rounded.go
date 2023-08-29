package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckroomRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20q-.725 0-.95-.688T2.4 18.2l8.6-6.45V10q0-.425.275-.713t.7-.287q.65.025 1.088-.413T13.5 7.5q0-.625-.425-1.063T12.025 6q-.45 0-.838.263t-.562.687q-.125.275-.363.413t-.537.137q-.55 0-.837-.463t-.063-.962q.425-.95 1.275-1.513T12 4q1.45 0 2.475 1T15.5 7.45q0 1.2-.7 2.125T13 10.85v.9l8.6 6.45q.575.425.35 1.113T21 20H3Zm3-2h12l-6-4.5L6 18Z"/>`),
		g.Group(children),
	)
}
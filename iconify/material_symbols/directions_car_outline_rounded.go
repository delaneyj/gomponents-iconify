package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsCarOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19v.5q0 .625-.438 1.063T4.5 21q-.625 0-1.063-.438T3 19.5v-7.15q0-.175.025-.35t.075-.325L4.975 6.35q.2-.6.725-.975T6.875 5h10.25q.65 0 1.175.375t.725.975l1.875 5.325q.05.15.075.325t.025.35v7.15q0 .625-.438 1.063T19.5 21q-.625 0-1.063-.438T18 19.5V19H6Zm-.2-9h12.4l-1.05-3H6.85L5.8 10ZM5 12v5v-5Zm2.5 4q.625 0 1.063-.438T9 14.5q0-.625-.438-1.063T7.5 13q-.625 0-1.063.438T6 14.5q0 .625.438 1.063T7.5 16Zm9 0q.625 0 1.063-.438T18 14.5q0-.625-.438-1.063T16.5 13q-.625 0-1.063.438T15 14.5q0 .625.438 1.063T16.5 16ZM5 17h14v-5H5v5Z"/>`),
		g.Group(children),
	)
}
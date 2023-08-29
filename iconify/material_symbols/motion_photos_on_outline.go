package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MotionPhotosOnOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-1.075.225-2.113t.65-2.012l1.55 1.55q-.2.65-.313 1.287T4 12q0 3.35 2.325 5.675T12 20q3.35 0 5.675-2.325T20 12q0-3.35-2.325-5.675T12 4q-.675 0-1.313.112t-1.262.313L7.9 2.9q1-.45 2-.675T12 2q2.075 0 3.9.787t3.175 2.138q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22ZM5.5 7q-.625 0-1.063-.438T4 5.5q0-.625.438-1.063T5.5 4q.625 0 1.063.438T7 5.5q0 .625-.438 1.063T5.5 7Zm6.5 5Z"/>`),
		g.Group(children),
	)
}
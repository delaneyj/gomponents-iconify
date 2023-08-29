package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MotionPhotosOnOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-.8.125-1.6T2.5 8.825q.125-.4.513-.537t.737.062q.375.2.537.588t.038.812q-.15.55-.238 1.113T4 12q0 3.35 2.325 5.675T12 20q3.35 0 5.675-2.325T20 12q0-3.35-2.325-5.675T12 4q-.6 0-1.188.087T9.65 4.35q-.425.125-.8-.025T8.3 3.8q-.175-.375-.013-.762t.563-.513q.75-.275 1.55-.4T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22ZM5.5 7q-.625 0-1.063-.438T4 5.5q0-.625.438-1.063T5.5 4q.625 0 1.063.438T7 5.5q0 .625-.438 1.063T5.5 7Zm6.5 5Z"/>`),
		g.Group(children),
	)
}
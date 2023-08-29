package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PsychiatryRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-.425 0-.713-.288T11 20v-6.975q-1.6 0-3.075-.613t-2.6-1.737Q4.2 9.55 3.6 8.075T3 5V4q0-.425.288-.712T4 3h1q1.575 0 3.05.613t2.6 1.737q.775.775 1.288 1.7t.787 1.975q.125-.175.275-.338t.325-.337q1.125-1.125 2.6-1.738T19 6h1q.425 0 .713.288T21 7v1q0 1.6-.613 3.075t-1.737 2.6Q17.525 14.8 16.063 15.4T13 16v4q0 .425-.288.713T12 21Z"/>`),
		g.Group(children),
	)
}
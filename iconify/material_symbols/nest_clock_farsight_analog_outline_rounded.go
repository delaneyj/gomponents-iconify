package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestClockFarsightAnalogOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13 12.175l2.25 2.25q.275.275.275.688t-.275.712q-.3.3-.713.3t-.712-.3L11.3 13.3q-.15-.15-.225-.338T11 12.575V9q0-.425.288-.713T12 8q.425 0 .713.288T13 9v3.175ZM12 6q-.425 0-.713-.288T11 5V4h2v1q0 .425-.288.713T12 6Zm6 6q0-.425.288-.713T19 11h1v2h-1q-.425 0-.713-.288T18 12Zm-6 6q.425 0 .713.288T13 19v1h-2v-1q0-.425.288-.713T12 18Zm-6-6q0 .425-.288.713T5 13H4v-2h1q.425 0 .713.288T6 12Zm6 10q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm8-10q0-3.35-2.325-5.675T12 4Q8.65 4 6.325 6.325T4 12q0 3.35 2.325 5.675T12 20q3.35 0 5.675-2.325T20 12Zm-8 0Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnsubscribeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v5.7q-.45-.225-.963-.375T20 11.1V8l-7.475 4.675q-.125.075-.263.113t-.262.037q-.125 0-.263-.037t-.262-.113L4 8v10h8q0 .525.075 1.012T12.3 20H4Zm8-9l8-5H4l8 5Zm-8 7v-7v.075V8v.25V6.8q0-.275 0 0V6v.8v-.013V8.25V8v10Zm15 5q-2.075 0-3.538-1.463T14 18q0-2.075 1.463-3.538T19 13q2.075 0 3.538 1.463T24 18q0 2.075-1.463 3.538T19 23Zm-2.5-4.5h5q.2 0 .35-.15T22 18q0-.2-.15-.35t-.35-.15h-5q-.2 0-.35.15T16 18q0 .2.15.35t.35.15Z"/>`),
		g.Group(children),
	)
}
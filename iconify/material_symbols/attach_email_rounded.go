package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttachEmailRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18q-.825 0-1.413-.588T1 16V4q0-.825.588-1.413T3 2h16q.825 0 1.413.588T21 4v6h-3.5q-1.45 0-2.475 1.025T14 13.5V18H3Zm8-9L4.3 4.8q-.425-.275-.863-.025T3 5.525q0 .225.1.413t.3.312l7.075 4.425q.25.15.525.15t.525-.15L18.6 6.25q.2-.125.3-.312t.1-.413q0-.5-.437-.75T17.7 4.8L11 9Zm8 13q-1.65 0-2.825-1.175T15 18v-4.5q0-1.05.725-1.775T17.5 11q1.05 0 1.775.725T20 13.5V17q0 .425-.288.713T19 18q-.425 0-.713-.288T18 17v-3.5q0-.2-.15-.35T17.5 13q-.2 0-.35.15t-.15.35V18q0 .825.588 1.413T19 20q.825 0 1.413-.588T21 18v-3q0-.425.288-.713T22 14q.425 0 .713.288T23 15v3q0 1.65-1.175 2.825T19 22Z"/>`),
		g.Group(children),
	)
}
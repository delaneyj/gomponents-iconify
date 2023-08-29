package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterLockSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 8h6V6q0-1.25-.875-2.125T12 3q-1.25 0-2.125.875T9 6v2Zm10 14q-1.05 0-1.775-.725T16.5 19.5q0-.5.175-.913t.475-.762L19 15.75l1.85 2.075q.3.35.475.763t.175.912q0 1.05-.725 1.775T19 22Zm-7-5q.825 0 1.413-.588T14 15q0-.825-.588-1.413T12 13q-.825 0-1.413.588T10 15q0 .825.588 1.413T12 17Zm-8 5V8h3V6q0-2.075 1.463-3.538T12 1q2.075 0 3.538 1.463T17 6v2h3v6.1q-.275-.05-.525-.075T18.95 14q-2.05 0-3.5 1.45T14 18.975q0 .825.25 1.6T15.025 22H4Z"/>`),
		g.Group(children),
	)
}
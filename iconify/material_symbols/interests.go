package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Interests(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2 11l5-9l5 9H2Zm5 10q-1.65 0-2.825-1.175T3 17q0-1.65 1.175-2.825T7 13q1.65 0 2.825 1.175T11 17q0 1.65-1.175 2.825T7 21Zm6 0v-8h8v8h-8Zm4-10q-1.425-1.2-2.388-2.025t-1.537-1.45Q12.5 6.9 12.25 6.35T12 5.175q0-1.125.788-1.9T14.75 2.5q.675 0 1.263.313t.987.862q.4-.55.988-.862T19.25 2.5q1.175 0 1.963.775t.787 1.9q0 .625-.25 1.175t-.825 1.175q-.575.625-1.538 1.45T17 11Z"/>`),
		g.Group(children),
	)
}
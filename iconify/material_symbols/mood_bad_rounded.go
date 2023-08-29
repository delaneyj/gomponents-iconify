package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodBadRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 11q.625 0 1.063-.438T17 9.5q0-.625-.438-1.063T15.5 8q-.625 0-1.063.438T14 9.5q0 .625.438 1.063T15.5 11Zm-7 0q.625 0 1.063-.438T10 9.5q0-.625-.438-1.063T8.5 8q-.625 0-1.063.438T7 9.5q0 .625.438 1.063T8.5 11ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm0-8.5q-1.425 0-2.675.7t-2 1.925q-.15.3.025.588t.525.287H16.1q.35 0 .525-.288t.025-.587Q15.9 14.9 14.662 14.2T12 13.5Z"/>`),
		g.Group(children),
	)
}
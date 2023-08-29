package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowseActivity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21v-2h22v2H1Zm3-3q-.825 0-1.413-.588T2 16v-5h5.375L9.1 14.45q.125.25.363.4t.512.15q.275 0 .525-.125t.375-.375l3.075-5.375l.65 1.325q.125.275.375.413T15.5 11H22v5q0 .825-.588 1.412T20 18H4Zm6.075-6.125L8.9 9.55q-.125-.25-.375-.4T8 9H2V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v4h-5.875L14.9 6.55q-.125-.275-.375-.413T14 6q-.275 0-.5.138t-.35.362l-3.075 5.375Z"/>`),
		g.Group(children),
	)
}
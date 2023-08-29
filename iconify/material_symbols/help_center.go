package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelpCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18q.525 0 .888-.363t.362-.887q0-.525-.363-.888T12 15.5q-.525 0-.888.363t-.362.887q0 .525.363.888T12 18Zm-.9-3.85h1.85q0-.9.2-1.325T14 11.75q.875-.875 1.237-1.463T15.6 8.95q0-1.325-.9-2.137T12.275 6q-1.375 0-2.337.675T8.6 8.55l1.65.65q.175-.675.7-1.087t1.225-.413q.675 0 1.125.363t.45.962q0 .425-.275.9t-.925 1.05q-.825.725-1.138 1.388T11.1 14.15ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}
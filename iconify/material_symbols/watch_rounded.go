package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 22q-.65 0-1.188-.388t-.737-1.037L7.65 17.45q-1.2-.95-1.925-2.375T5 12q0-1.65.725-3.075T7.65 6.55l.925-3.125q.2-.65.738-1.037T10.5 2h3.375q.5 0 .9.288t.55.787L16.35 6.55q1.2.95 1.925 2.375T19 12q0 1.65-.725 3.075T16.35 17.45l-.925 3.125q-.2.65-.738 1.038T13.5 22h-3Zm1.5-5q2.075 0 3.538-1.463T17 12q0-2.075-1.463-3.538T12 7Q9.925 7 8.462 8.463T7 12q0 2.075 1.463 3.538T12 17Z"/>`),
		g.Group(children),
	)
}
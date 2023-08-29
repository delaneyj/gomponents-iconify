package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SavedSearchRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.5 10.95l-1.375 1.075q-.15.125-.3.013t-.1-.288l.525-1.7L6.8 8.9q-.125-.125-.063-.288t.238-.162H8.7l.55-1.725q.05-.175.25-.175t.25.175l.55 1.725h1.725q.175 0 .238.163T12.2 8.9l-1.45 1.15l.525 1.7q.05.175-.1.288t-.3-.013L9.5 10.95Zm0 5.05q-2.725 0-4.612-1.888T3 9.5q0-2.725 1.888-4.612T9.5 3q2.725 0 4.612 1.888T16 9.5q0 1.1-.35 2.075T14.7 13.3l5.6 5.6q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-5.6-5.6q-.75.6-1.725.95T9.5 16Zm0-2q1.875 0 3.188-1.313T14 9.5q0-1.875-1.313-3.188T9.5 5Q7.625 5 6.312 6.313T5 9.5q0 1.875 1.313 3.188T9.5 14Z"/>`),
		g.Group(children),
	)
}
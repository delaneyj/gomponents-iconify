package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SentimentVerySatisfied(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17.5q1.7 0 3.088-.963T17.1 14H6.9q.625 1.575 2.013 2.538T12 17.5ZM7.8 11l1.1-1.05L9.95 11L11 9.95L8.9 7.8L6.75 9.95L7.8 11Zm6.25 0l1.05-1.05L16.2 11l1.05-1.05L15.1 7.8L13 9.95L14.05 11ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}
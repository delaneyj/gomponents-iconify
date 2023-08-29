package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GasStation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M8.5 21V3.5l.088-.059a11.56 11.56 0 0 1 12.824 0l.088.059V21M6 22.5h18M0 4.5h4.5v8a4 4 0 0 0 4 4m-7-12a3 3 0 0 0 3 3m14 2h-7V5.043a11.559 11.559 0 0 1 7 0V9.5Z"/>`),
		g.Group(children),
	)
}
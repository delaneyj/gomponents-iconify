package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatureEcologyPottedTreeTwoTreePlantPot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.25 13.5h-4l-.5-4h5l-.5 4zm2.88-7.43a2 2 0 1 1-2.47-3.14c.86-.68 3.59-.28 3.59-.28S13 5.39 12.13 6.07Z"/><path d="M9.25 5.64a5.5 5.5 0 0 0-2 3.86c0-3-1.5-5-4-6.5"/><path d="M1.67 5a2.56 2.56 0 0 0 3.61-3.61C4.28.4.75.5.75.5S.67 4 1.67 5Z"/></g>`),
		g.Group(children),
	)
}
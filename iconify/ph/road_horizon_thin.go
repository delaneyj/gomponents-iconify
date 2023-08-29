package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadHorizonThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M235.49 190a4 4 0 0 1-1.53 5.45a4.07 4.07 0 0 1-2 .51a4 4 0 0 1-3.49-2L157.66 68H132v12a4 4 0 0 1-8 0V68H98.34L27.49 194a4 4 0 0 1-3.49 2a4.07 4.07 0 0 1-2-.51a4 4 0 0 1-1.49-5.49L89.16 68H24a4 4 0 0 1 0-8h208a4 4 0 0 1 0 8h-65.16ZM128 116a4 4 0 0 0-4 4v16a4 4 0 0 0 8 0v-16a4 4 0 0 0-4-4Zm0 56a4 4 0 0 0-4 4v16a4 4 0 0 0 8 0v-16a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}
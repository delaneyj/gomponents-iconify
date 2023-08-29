package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PencilLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.728 9.576l-1.414-1.414L5 17.476v1.414h1.414l9.314-9.314Zm1.414-1.414l1.414-1.414l-1.414-1.414l-1.414 1.414l1.414 1.414Zm-9.9 12.728H3v-4.243L16.435 3.212a1 1 0 0 1 1.414 0l2.829 2.829a1 1 0 0 1 0 1.414L7.243 20.89Z"/>`),
		g.Group(children),
	)
}
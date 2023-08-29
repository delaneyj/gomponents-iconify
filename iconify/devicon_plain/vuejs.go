package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vuejs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="m0 8.934l49.854.158l14.3 24.415l14.3-24.415l49.548-.158l-63.835 110.134zm126.987.637l-24.37.021l-38.473 66.053L25.692 9.592l-24.75-.02l63.212 107.89z"/>`),
		g.Group(children),
	)
}
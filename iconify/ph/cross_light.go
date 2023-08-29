package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 74h-40a2 2 0 0 1-2-2V32a14 14 0 0 0-14-14h-32a14 14 0 0 0-14 14v40a2 2 0 0 1-2 2H56a14 14 0 0 0-14 14v32a14 14 0 0 0 14 14h40a2 2 0 0 1 2 2v88a14 14 0 0 0 14 14h32a14 14 0 0 0 14-14v-88a2 2 0 0 1 2-2h40a14 14 0 0 0 14-14V88a14 14 0 0 0-14-14Zm2 46a2 2 0 0 1-2 2h-40a14 14 0 0 0-14 14v88a2 2 0 0 1-2 2h-32a2 2 0 0 1-2-2v-88a14 14 0 0 0-14-14H56a2 2 0 0 1-2-2V88a2 2 0 0 1 2-2h40a14 14 0 0 0 14-14V32a2 2 0 0 1 2-2h32a2 2 0 0 1 2 2v40a14 14 0 0 0 14 14h40a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}
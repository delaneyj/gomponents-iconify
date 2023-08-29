package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.427 1.963l5.61 5.61L7.61 22.001H2v-5.61L16.427 1.962Zm0 2.828l-2.781 2.782l2.781 2.782l2.782-2.782l-2.782-2.782Zm-1.414 6.978l-2.782-2.782L4 17.22V20h2.782l8.231-8.232Zm7.212 10.232h-9.543v-2h9.543v2Z"/>`),
		g.Group(children),
	)
}
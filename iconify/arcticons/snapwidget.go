package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snapwidget(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v16.559a6.387 6.387 0 0 1 1.404-.176h.006a6.4 6.4 0 0 1 5.303 2.816a6.414 6.414 0 0 1 8.89-1.803q.15.099.293.206a6.45 6.45 0 0 1 1.11 9.277l-.05.064l-.052.057l-6.799 8H40.5a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.417 30.354a2.246 2.246 0 0 0-1.745 3.664l3.603 4.229l3.569-4.186l.017-.02l.018-.023a2.25 2.25 0 1 0-3.604-2.68a2.243 2.243 0 0 0-1.856-.984Z"/>`),
		g.Group(children),
	)
}
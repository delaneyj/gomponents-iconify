package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineglassTriangleBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M19.3 3H4.7C3.2 3 2.439 4.794 3.485 5.864l7.801 7.976a1 1 0 0 0 1.43 0l7.801-7.976C21.562 4.794 20.8 3 19.3 3Z" opacity=".5"/><path fill-rule="evenodd" d="M7.006 21a.75.75 0 0 1 .75-.75h8.488a.75.75 0 0 1 0 1.5H7.756a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/><path d="M11.285 13.84a1 1 0 0 0 1.43 0L16.471 10H7.53l3.756 3.84Z"/><path d="M11.285 13.84a1 1 0 0 0 1.43 0l.035-.035v6.445h-1.5v-6.446l.035.037Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}
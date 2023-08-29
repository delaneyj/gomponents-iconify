package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CriticalCare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M9 6a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H9Zm4 7a1 1 0 0 0-1 1v15.96a1 1 0 0 0 1 1h10v2.142h-4.538v2h11.076v-2H25v-2.143h10a1 1 0 0 0 1-1V14a1 1 0 0 0-1-1H13Z" clip-rule="evenodd"/><path d="M13.846 21.679h5.658l3.403-5.985l2.127 7.363l2.602-3.31h5.513a1 1 0 0 1 1.005.994a1 1 0 0 1-1.005.993h-4.53l-4.43 5.633l-1.891-6.548l-1.62 2.847h-6.832V21.68Z"/></g>`),
		g.Group(children),
	)
}
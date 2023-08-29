package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlaskOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 3h6m-2 6h1m-4-6v3m-.268 3.736L6 20a.7.7 0 0 0 .5 1h11a.7.7 0 0 0 .5-1l-1.143-3.142m-2.288-6.294L14 9V3M3 3l18 18"/>`),
		g.Group(children),
	)
}
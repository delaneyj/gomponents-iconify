package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Carrot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M15.624 20.682C14.29 15.248 18.404 10 24 10s9.71 5.248 8.376 10.682L27.279 41.43a3.376 3.376 0 0 1-6.557 0l-5.098-20.747Z"/><path stroke-linecap="round" d="M24 4v5.5m6.102-3.908l-2.728 3.25M18 5.592l2.727 3.25M16 19h6m3 7h6m-12 8h4"/></g>`),
		g.Group(children),
	)
}
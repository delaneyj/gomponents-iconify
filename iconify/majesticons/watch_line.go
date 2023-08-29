package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 9v3l2 1m4-1a6 6 0 1 1-12 0a6 6 0 0 1 12 0z"/><path d="M15 6.5V3a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v3.5m6 11V21a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-3.5"/></g>`),
		g.Group(children),
	)
}
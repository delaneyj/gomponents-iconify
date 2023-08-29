package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4.5 3.5h6m-6 0a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1m0-8v-2a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2m0 0a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1m0 0h-6m6 0v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2m9-5.5v3"/>`),
		g.Group(children),
	)
}
package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 16.5v-3a1 1 0 1 1 2 0v3a1 1 0 0 1-2 0zm4 0v-6a1 1 0 1 1 2 0v6a1 1 0 0 1-2 0zm4 0v-9a1 1 0 1 1 2 0v9a1 1 0 0 1-2 0z"/>`),
		g.Group(children),
	)
}
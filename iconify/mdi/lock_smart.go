package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockSmart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a6 6 0 0 0-6 6v8a6 6 0 0 0 6 6a6 6 0 0 0 6-6V8a6 6 0 0 0-6-6M8 6h2v2H8V6m3 0h2v2h-2V6m3 0h2v2h-2V6M8 9h2v2H8V9m3 0h2v2h-2V9m3 0h2v2h-2V9m-6 3h2v2H8v-2m3 0h2v2h-2v-2m3 0h2v2h-2v-2m-2 4a2 2 0 0 1 2 2a2 2 0 0 1-2 2a2 2 0 0 1-2-2a2 2 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}
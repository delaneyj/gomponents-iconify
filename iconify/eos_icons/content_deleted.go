package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentDeleted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 19H5V5h6V3H5a2.006 2.006 0 0 0-2 2v14a2.006 2.006 0 0 0 2 2h14a2.006 2.006 0 0 0 2-2v-4h-2Z"/><path fill="currentColor" d="M15 5h6v6a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1V5Zm7-2h-2l-.571-1h-2.858L16 3h-2v1h8V3z"/>`),
		g.Group(children),
	)
}
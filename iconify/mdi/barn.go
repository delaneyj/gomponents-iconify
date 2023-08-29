package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3L3 8.2V21h6l2.9-3l3.1 3h6V8.2L12 3M7.9 20v-6l3 3l-3 3m1-7h6l-3 3l-3-3m7 7l-3-3l3-3v6m-.9-9H8.8V9H15v2Z"/>`),
		g.Group(children),
	)
}
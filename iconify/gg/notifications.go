package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notifications(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M20 7a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/><path d="M12 6H4v14h14v-8h-2v6H6V8h6V6Z"/></g>`),
		g.Group(children),
	)
}
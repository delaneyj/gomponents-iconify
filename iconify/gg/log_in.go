package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M15.486 20h4a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-4v2h4v12h-4v2Z"/><path d="m10.158 17.385l-1.42-1.408l3.92-3.953H3.513a1 1 0 1 1 0-2h9.163l-3.98-3.947l1.408-1.42l6.391 6.337l-6.337 6.391Z"/></g>`),
		g.Group(children),
	)
}
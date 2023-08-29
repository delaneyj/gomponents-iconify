package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soccer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M11 1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0zm0 9.5a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm1.84-4.91l-1.91-1.91a.48.48 0 0 0-.37-.18H3.5a.5.5 0 0 0 0 1h2.7L3 11.3a.488.488 0 0 0 0 .2a.511.511 0 0 0 1 .21L5 10h2l-1.93 4.24a.49.49 0 0 0-.07.26a.51.51 0 0 0 1 .2l4.7-9.38l1.44 1.48a.5.5 0 0 0 .7-.71z"/>`),
		g.Group(children),
	)
}
package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 1c-2 0-7 .25-6.5.75L7 8v4c0 1-3 .5-3 2h7c0-1.5-3-1-3-2V8l6-6.25c.5-.5-4.5-.75-6.5-.75zm0 1c2.5 0 4.75.25 4.75.25L11.5 3h-8l-.75-.75S5 2 7.5 2z"/>`),
		g.Group(children),
	)
}
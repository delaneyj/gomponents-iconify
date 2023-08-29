package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Overflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M22 11a10 10 0 0 1-20 0h20Z" opacity=".2"/><path d="M20 11a8 8 0 0 1-16 0h16Z" opacity=".3"/><path d="M20 11a8 8 0 0 0-16 0h16Z"/></g>`),
		g.Group(children),
	)
}
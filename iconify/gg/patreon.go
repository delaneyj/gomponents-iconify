package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Patreon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M21 10a6 6 0 1 1-12 0a6 6 0 0 1 12 0Z" opacity=".5"/><path d="M3 4h4v16H3V4Z"/></g>`),
		g.Group(children),
	)
}
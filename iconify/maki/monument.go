package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 0L6 2.5v7h3v-7L7.5 0zM3 11.5V15h9v-3.5L10.5 10h-6L3 11.5z"/>`),
		g.Group(children),
	)
}
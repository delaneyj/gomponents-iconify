package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flutter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19.083 0L3.068 16L8 20.932L28.912.016h-9.808zm.021 14.76l-8.631 8.609L19.104 32h9.828l-8.615-8.625l8.615-8.615z"/>`),
		g.Group(children),
	)
}
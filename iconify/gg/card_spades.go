package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardSpades(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9.146 11.707a2 2 0 0 0 2.829 2.829L12 14.51l.025.026a2 2 0 1 0 2.829-2.829l-2.83-2.827l-.024.024l-.025-.025l-2.829 2.828Z"/><path fill-rule="evenodd" d="M3 20a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3H6a3 3 0 0 0-3 3v16Zm3 1h12a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
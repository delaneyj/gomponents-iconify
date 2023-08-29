package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 11a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path fill-rule="evenodd" d="M18 9a6.002 6.002 0 0 1-5 5.917V20a1 1 0 1 1-2 0v-5.083A6.002 6.002 0 0 1 12 3a6 6 0 0 1 6 6Zm-6 4a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
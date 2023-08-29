package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 14a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/><path fill-rule="evenodd" d="M8 5a1 1 0 0 1 0-2h8a1 1 0 1 1 0 2h-3v5.083A6.002 6.002 0 0 1 12 22a6 6 0 0 1-1-11.917V5H8Zm4 7a4 4 0 1 1 0 8a4 4 0 0 1 0-8Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
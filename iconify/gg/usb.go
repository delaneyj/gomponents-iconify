package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10 4.5h1v2h-1v-2Zm4 0h-1v2h1v-2Z"/><path fill-rule="evenodd" d="M7 8.5v-7h10v7h2v11a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3v-11h2Zm2-5h6v5H9v-5Zm8 7H7v9a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-9Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
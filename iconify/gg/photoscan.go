package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Photoscan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path fill-rule="evenodd" d="M17 3a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h10Zm-4.535 2H17v11H7v-5.535A4 4 0 0 0 12.465 5ZM9 5a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
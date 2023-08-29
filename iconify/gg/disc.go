package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm-1 3a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z" clip-rule="evenodd"/><path d="M5 12a7 7 0 0 1 7-7v2a5 5 0 0 0-5 5H5Zm7 5a5 5 0 0 0 5-5h2a7 7 0 0 1-7 7v-2Z"/><path fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1ZM3 12a9 9 0 1 0 18 0a9 9 0 0 0-18 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
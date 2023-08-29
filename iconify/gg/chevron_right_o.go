package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRightO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11.086 7.757L15.328 12l-4.242 4.243l-1.414-1.414L12.5 12L9.672 9.172l1.414-1.415Z"/><path fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1Zm9 11a9 9 0 1 0-18 0a9 9 0 0 0 18 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
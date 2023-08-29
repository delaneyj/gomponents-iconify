package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeftO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m12 7.757l1.414 1.415L10.586 12l2.828 2.829L12 16.242L7.757 12L12 7.757Z"/><path fill-rule="evenodd" d="M12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1ZM3 12a9 9 0 1 1 18 0a9 9 0 0 1-18 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
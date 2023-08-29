package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushChevronDownO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16 14v2H8v-2h8ZM7.757 8.703l1.415-1.415L12 10.117l2.828-2.829l1.415 1.415L12 12.945L7.757 8.703Z"/><path fill-rule="evenodd" d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11Zm0-2a9 9 0 1 0 0-18a9 9 0 0 0 0 18Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
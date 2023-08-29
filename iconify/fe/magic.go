package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feMagic0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMagic1" fill="currentColor"><path id="feMagic2" d="m3 5l2-2l16 16l-2 2L3 5Zm10 0l1-2l1 2l2 1l-2 1l-1 2l-1-2l-2-1l2-1ZM5 15l1-2l1 2l2 1l-2 1l-1 2l-1-2l-2-1l2-1ZM4 9l1 1l-1 1l-1-1l1-1Zm14 1l1 1l-1 1l-1-1l1-1Z"/></g></g>`),
		g.Group(children),
	)
}
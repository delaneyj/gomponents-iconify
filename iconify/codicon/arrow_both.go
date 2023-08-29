package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBoth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m3 9l2.146 2.146l-.707.708l-3-3v-.708l3-3l.707.708L3 8h10l-2.146-2.146l.707-.708l3 3v.708l-3 3l-.707-.707L13 9H3z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
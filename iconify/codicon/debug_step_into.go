package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DebugStepInto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 9.532h.542l3.905-3.905l-1.061-1.06l-2.637 2.61V1H7.251v6.177l-2.637-2.61l-1.061 1.06l3.905 3.905H8zm1.956 3.481a2 2 0 1 1-4 0a2 2 0 0 1 4 0z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
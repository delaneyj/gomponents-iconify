package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DebugStop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m13 2l1 1v10l-1 1H3l-1-1V3l1-1h10Zm-.254 1.25H3.255v9.5h9.491v-9.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
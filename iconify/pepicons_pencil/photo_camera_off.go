package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoCameraOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.172 5.232L5.762 6.5H4.5A2.5 2.5 0 0 0 2 9v6a2.5 2.5 0 0 0 2.5 2.5h11A2.5 2.5 0 0 0 18 15V9a2.5 2.5 0 0 0-2.5-2.5h-1.263l-.409-1.268a2.5 2.5 0 0 0-2.38-1.732H8.552a2.5 2.5 0 0 0-2.38 1.732ZM4.5 7.5h1.99l.633-1.96A1.5 1.5 0 0 1 8.551 4.5h2.898a1.5 1.5 0 0 1 1.427 1.04l.633 1.96H15.5A1.5 1.5 0 0 1 17 9v6a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 3 15V9a1.5 1.5 0 0 1 1.5-1.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7 11.5a3 3 0 1 0 6 0a3 3 0 0 0-6 0Zm5 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}
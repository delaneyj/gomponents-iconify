package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestartBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.873 7.873a7.253 7.253 0 0 1 7.928-1.563l1.13-1.13a8.75 8.75 0 1 0 4.751 6.727a.75.75 0 0 0-1.488.187a7.25 7.25 0 1 1-12.32-4.22Z" clip-rule="evenodd" opacity=".5"/><path d="M18.721 4.201a.75.75 0 0 0-1.28-.53l-1.51 1.51l-1.13 1.13l-1.603 1.603a.75.75 0 0 0 .53 1.28h4.243a.75.75 0 0 0 .75-.75V4.2Z"/></g>`),
		g.Group(children),
	)
}
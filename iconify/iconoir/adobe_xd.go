package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeXd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7v10a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4V7a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4ZM7 8l4 8m-4 0l4-8"/><path d="M17 12v3.4a.6.6 0 0 1-.6.6H15a2 2 0 0 1-2-2v0a2 2 0 0 1 2-2h2Zm0 0V9"/></g>`),
		g.Group(children),
	)
}
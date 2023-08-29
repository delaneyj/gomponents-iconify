package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Link(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feLink0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLink1" fill="currentColor"><path id="feLink2" d="M11.874 11h.252c.444-1.725 2.01-3 3.874-3h2a4 4 0 1 1 0 8h-2a4.002 4.002 0 0 1-3.874-3h-.252A4.002 4.002 0 0 1 8 16H6a4 4 0 1 1 0-8h2a4.002 4.002 0 0 1 3.874 3Zm-2.124.031A2 2 0 0 0 8 10H6a2 2 0 1 0 0 4h2a2 2 0 0 0 1.75-1.031a1 1 0 0 1 0-1.938ZM14 12.97A2 2 0 0 0 15.75 14h2a2 2 0 1 0 0-4h-2A2 2 0 0 0 14 11.031a1 1 0 0 1 0 1.938Z"/></g></g>`),
		g.Group(children),
	)
}
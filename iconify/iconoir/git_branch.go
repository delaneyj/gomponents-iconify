package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitBranch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM6 20a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0-4V3"/><path d="M8 18h1c3.5 0 9-2.1 9-8.5V8"/></g>`),
		g.Group(children),
	)
}
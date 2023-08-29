package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitPullRequest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 21a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM6 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM6 7v10m12 0V7s0-2-2-2h-3"/><path d="M15 7.5L12.5 5L15 2.5"/></g>`),
		g.Group(children),
	)
}
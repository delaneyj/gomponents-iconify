package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodcastLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M9 10a3 3 0 1 1 6 0v3a3 3 0 1 1-6 0v-3Z"/><path stroke-linecap="round" d="M13 10h2m-2 3h2m-6-3h1m-1 3h1" opacity=".5"/><path d="M4.154 16C5.174 16 6 15.173 6 14.154V10a6 6 0 1 1 12 0v4.154c0 1.02.826 1.846 1.846 1.846" opacity=".5"/><path d="M2 12a2 2 0 1 1 4 0v2a2 2 0 1 1-4 0v-2Zm16 0a2 2 0 1 1 4 0v2a2 2 0 1 1-4 0v-2Z"/><path stroke-linecap="round" d="M12 16v3" opacity=".5"/></g>`),
		g.Group(children),
	)
}
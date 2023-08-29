package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodcastBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M9 10a3 3 0 1 1 6 0v3a3 3 0 1 1-6 0v-3Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M13 10h2m-2 3h2m-6-3h1m-1 3h1"/><path fill="currentColor" d="M16.951 8.25a.75.75 0 1 0 1.415-.5l-1.415.5ZM6.75 14.154V10h-1.5v4.154h1.5Zm-1.5 0c0 .605-.49 1.096-1.096 1.096v1.5a2.596 2.596 0 0 0 2.596-2.596h-1.5Zm12 0a2.596 2.596 0 0 0 2.596 2.596v-1.5c-.605 0-1.096-.49-1.096-1.096h-1.5ZM12 3.25A6.75 6.75 0 0 0 5.25 10h1.5c0-2.9 2.35-5.25 5.25-5.25v-1.5Zm0 1.5c2.285 0 4.23 1.46 4.951 3.5l1.415-.5A6.752 6.752 0 0 0 12 3.25v1.5Z"/><path stroke="currentColor" stroke-width="1.5" d="M2 12a2 2 0 1 1 4 0v2a2 2 0 1 1-4 0v-2Zm16 0a2 2 0 1 1 4 0v2a2 2 0 1 1-4 0v-2Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 16v3"/></g>`),
		g.Group(children),
	)
}
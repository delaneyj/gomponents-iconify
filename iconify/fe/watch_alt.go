package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feWatchAlt0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feWatchAlt1" fill="currentColor"><path id="feWatchAlt2" d="M15 6h1a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-1v2a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-2H8a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h1V4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2ZM8 8v8h8V8H8Z"/></g></g>`),
		g.Group(children),
	)
}
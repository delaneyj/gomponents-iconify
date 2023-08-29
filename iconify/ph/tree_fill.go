package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M128 187.88h-.06a71.3 71.3 0 0 1-8 4.61V232a8 8 0 0 0 16 0v-39.51a71.3 71.3 0 0 1-8-4.61Z"/><path d="M198.1 62.6a76 76 0 0 0-140.2 0A72.29 72.29 0 0 0 16 127.8c-.11 38.82 31.36 71.2 70.14 72.2a71.68 71.68 0 0 0 33.86-7.51v-35.55l-43.58-21.78a8 8 0 1 1 7.16-14.32L120 139.06V88a8 8 0 0 1 16 0v27.06l36.42-18.22a8 8 0 1 1 7.16 14.32L136 132.94v59.55a71.45 71.45 0 0 0 32 7.51h1.86c38.78-1 70.25-33.36 70.14-72.18a72.26 72.26 0 0 0-41.9-65.22Z"/></g>`),
		g.Group(children),
	)
}
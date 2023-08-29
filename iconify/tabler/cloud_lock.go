package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 18a3.5 3.5 0 0 0 0-7h-1c.397-1.768-.285-3.593-1.788-4.787c-1.503-1.193-3.6-1.575-5.5-1S7.397 7.232 7 9c-2.199-.088-4.155 1.326-4.666 3.373c-.512 2.047.564 4.154 2.566 5.027"/><path d="M8 16a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1zm2-1v-2a2 2 0 1 1 4 0v2"/></g>`),
		g.Group(children),
	)
}
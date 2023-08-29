package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderHermaphrodite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 14v7m-3-3h6M12 6a4 4 0 1 1 0 8a4 4 0 0 1 0-8z"/><path d="M15 3a3 3 0 1 1-6 0"/></g>`),
		g.Group(children),
	)
}
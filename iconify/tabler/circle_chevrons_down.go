package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleChevronsDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 9l-3 3l-3-3m6 4l-3 3l-3-3"/><path d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18z"/></g>`),
		g.Group(children),
	)
}
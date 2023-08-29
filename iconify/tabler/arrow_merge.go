package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMerge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 7l4-4l4 4"/><path d="M12 3v5.394A6.737 6.737 0 0 1 9 14a6.737 6.737 0 0 0-3 5.606V21"/><path d="M12 3v5.394A6.737 6.737 0 0 0 15 14a6.737 6.737 0 0 1 3 5.606V21"/></g>`),
		g.Group(children),
	)
}
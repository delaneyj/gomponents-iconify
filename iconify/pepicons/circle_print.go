package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclePrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><circle cx="12" cy="11" r="6" opacity=".8"/><path fill-rule="evenodd" d="M4 10a6 6 0 1 0 12 0a6 6 0 0 0-12 0Zm11 0a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoCirclePrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><circle cx="12" cy="11" r="7" opacity=".8"/><path d="M9.5 9a.5.5 0 0 1 1 0v5a.5.5 0 0 1-1 0V9Z"/><path fill-rule="evenodd" d="M3 10a7 7 0 1 0 14 0a7 7 0 0 0-14 0Zm13 0a6 6 0 1 1-12 0a6 6 0 0 1 12 0Z" clip-rule="evenodd"/><circle cx="10" cy="6.75" r=".75"/></g>`),
		g.Group(children),
	)
}
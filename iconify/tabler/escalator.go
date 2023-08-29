package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Escalator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.5 5h-2.672a2 2 0 0 0-1.414.586L7 14H4.5a2.5 2.5 0 1 0 0 5h3.672a2 2 0 0 0 1.414-.586L18 10h1.5a2.5 2.5 0 0 0 0-5z"/>`),
		g.Group(children),
	)
}
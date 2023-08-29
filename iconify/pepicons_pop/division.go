package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Division(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M5 11a1 1 0 1 1 0-2h10a1 1 0 1 1 0 2H5Z"/><circle cx="10" cy="5.5" r="1.5"/><circle cx="10" cy="14.5" r="1.5"/></g>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 5h4V3a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v2z"/><path d="M14 3.5c0 1.626.507 3.212 1.45 4.537l.05.07a8.093 8.093 0 0 1 1.5 4.694V19a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-6.2c0-1.682.524-3.322 1.5-4.693l.05-.07A7.823 7.823 0 0 0 10 3.5"/><path d="M7 14.803A2.4 2.4 0 0 0 8 14a2.4 2.4 0 0 1 2-1a2.4 2.4 0 0 1 2 1a2.4 2.4 0 0 0 2 1a2.4 2.4 0 0 0 2-1a2.4 2.4 0 0 1 1-.805"/></g>`),
		g.Group(children),
	)
}
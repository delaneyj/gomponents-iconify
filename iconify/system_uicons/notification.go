package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notification(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 2)"><path d="M14.5 6.5v7a2 2 0 0 1-2 2H2.543a2 2 0 0 1-2-1.991l-.043-10A2 2 0 0 1 2.49 1.5H9.5"/><circle cx="14" cy="2" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
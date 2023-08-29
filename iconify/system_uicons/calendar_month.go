package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarMonth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 2.5h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2v-12a2 2 0 0 1 2-2zm-2 4h16"/><g fill="currentColor" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="1"/><circle cx="4.5" cy="8.5" r="1"/><circle cx="12.5" cy="8.5" r="1"/><circle cx="8.5" cy="12.5" r="1"/><circle cx="4.5" cy="12.5" r="1"/><circle cx="12.5" cy="12.5" r="1"/></g></g>`),
		g.Group(children),
	)
}
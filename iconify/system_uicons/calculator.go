package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 4.5h6a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-9a2 2 0 0 1 2-2zm-2 5h10"/><g fill="currentColor" transform="translate(5 4)"><circle cx="2.5" cy="7.5" r="1"/><circle cx="4.5" cy="7.5" r="1"/><circle cx="6.5" cy="7.5" r="1"/><circle cx="8.5" cy="7.5" r="1"/><circle cx="2.5" cy="9.5" r="1"/><circle cx="4.5" cy="9.5" r="1"/><circle cx="6.5" cy="9.5" r="1"/><circle cx="8.5" cy="9.5" r="1"/><circle cx="2.5" cy="11.5" r="1"/><circle cx="4.5" cy="11.5" r="1"/><circle cx="6.5" cy="11.5" r="1"/><circle cx="8.5" cy="11.5" r="1"/></g></g>`),
		g.Group(children),
	)
}
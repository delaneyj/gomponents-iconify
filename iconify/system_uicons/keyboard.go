package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.5 13.5v-6a2 2 0 0 0-2-2h-14a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2z"/><g fill="currentColor" transform="translate(1 5)"><circle cx="3.5" cy="2.5" r="1"/><circle cx="6.5" cy="2.5" r="1"/><circle cx="9.5" cy="2.5" r="1"/><circle cx="12.5" cy="2.5" r="1"/><circle cx="15.5" cy="2.5" r="1"/><circle cx="3.5" cy="4.5" r="1"/><circle cx="6.5" cy="4.5" r="1"/><circle cx="9.5" cy="4.5" r="1"/><circle cx="12.5" cy="4.5" r="1"/><circle cx="15.5" cy="4.5" r="1"/><circle cx="3.5" cy="6.5" r="1"/><circle cx="6.5" cy="6.5" r="1"/><circle cx="9.5" cy="6.5" r="1"/><circle cx="12.5" cy="6.5" r="1"/><circle cx="15.5" cy="6.5" r="1"/><circle cx="3.5" cy="8.5" r="1"/><circle cx="15.5" cy="8.5" r="1"/></g><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 13.5h6"/></g>`),
		g.Group(children),
	)
}
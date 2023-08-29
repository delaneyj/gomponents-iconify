package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FdroidNethunter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.6 29.4v8m-13.2 0v-8l5.3 8v-8m2.6 0v8m0-4h5.3"/><circle cx="17" cy="15.3" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31" cy="15.3" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.2 8.8h25.7a2.862 2.862 0 0 1 2.8 2.8V19a2.862 2.862 0 0 1-2.8 2.8H11.2A2.862 2.862 0 0 1 8.4 19v-7.4a2.734 2.734 0 0 1 2.8-2.8ZM5.5 5.9l3.7 3.7m2 15.2h25.7a2.862 2.862 0 0 1 2.8 2.8v11.7a2.862 2.862 0 0 1-2.8 2.8H11.2a2.862 2.862 0 0 1-2.8-2.8V27.6a2.734 2.734 0 0 1 2.8-2.8ZM42.5 5.9l-3.7 3.7"/>`),
		g.Group(children),
	)
}
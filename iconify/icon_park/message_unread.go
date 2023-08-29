package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageUnread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 16V36H29L24 41L19 36H4V6H34"/><path stroke="#000" stroke-linecap="round" stroke-width="4" d="M23 20H25.0025"/><path stroke="#000" stroke-linecap="round" stroke-width="4" d="M33.001 20H34.9999"/><path stroke="#000" stroke-linecap="round" stroke-width="4" d="M13.001 20H14.9999"/><circle cx="43" cy="7" r="3" fill="#000"/></g>`),
		g.Group(children),
	)
}
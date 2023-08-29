package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-width="4" d="M10 44H38C39.1046 44 40 43.1046 40 42V14.8847C40 14.3212 39.7623 13.7839 39.3453 13.4049L29.5721 4.52012C29.204 4.18544 28.7243 4 28.2268 4H10C8.89543 4 8 4.89543 8 6V42C8 43.1046 8.89543 44 10 44Z"/><circle cx="17" cy="14" r="3" fill="#000"/><rect width="16" height="14" x="16" y="23" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/></g>`),
		g.Group(children),
	)
}
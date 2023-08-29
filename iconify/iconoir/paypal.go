package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paypal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 17.5L6 3h7c6 0 6 9 0 9H8.7l-1.2 5.5H3Z"/><path d="m6.8 21l3-14.5h7c6 0 6 9 0 9h-4.3L11.3 21H6.8Z"/></g>`),
		g.Group(children),
	)
}
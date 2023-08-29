package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BestBefore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.4 19.8h-7.9l-7.8-13c-.5-1-1.8-1.3-2.7-.8c-.3.2-.6.5-.8.8l-7.8 13H6.6c-1.7 0-3.1 1.4-3.1 3.1c0 .3 0 .5.1.8L8 40.1c.3 1.2 1.5 2.1 2.7 2.1h26.5c1.3 0 2.4-.9 2.8-2.1l4.4-16.3c.4-1.7-.5-3.4-2.2-3.8c-.3-.1-.5-.2-.8-.2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 11.8l4.8 8h-9.6l4.8-8z"/><circle cx="24" cy="31" r="4.8" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
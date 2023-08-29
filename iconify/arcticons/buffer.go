package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buffer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.53 13.5L24 4.5l-16.53 9L24 22.49l16.53-8.99z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.12 18.75L7.47 24L24 33l16.53-9l-9.65-5.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.12 29.25L7.47 34.5l16.53 9l16.53-9l-9.65-5.25"/>`),
		g.Group(children),
	)
}
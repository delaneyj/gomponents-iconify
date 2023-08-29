package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlltidOppet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 8.813h-.1c-4.7-4.7-12.3-4.7-16.9 0c-4.7 4.7-4.7 12.3 0 16.9h0l17 17l17-17c4.7-4.7 4.7-12.3 0-16.9s-12.4-4.7-17 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.7 25.012h-3c-1.6 0-2.8-1.3-2.8-2.8v-10"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.7 27.712h-3.1c-2.9 0-5.3-2.4-5.3-5.3v-10.1m-11.2 12.7h3c1.6 0 2.8-1.3 2.8-2.8v-10"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.1 27.712h3.1c2.9 0 5.3-2.4 5.3-5.3v-10.1"/>`),
		g.Group(children),
	)
}
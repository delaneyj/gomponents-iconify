package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NormalApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.5 20.8c3.3.6 7.4 1.7 6.9 3.9c-.4 2.2-4.8 1.7-8 .9m0 0c.1.7.3 1.3.6 1.9c.3.3.5.6.9.9c-.4.3-.9.6-1.4.8m0 0s.3 2-.3 2.7c-.5.6-1.6.4-2.3.8c-.3.2-.8.5-.9.9c-.1.3.3.9.3.9l.7 1.9c1.2.7 2.3 1.6 3.4 2.5s1.9 2.1 2.3 3.5m-2.2-21.6s.3-5.8-1.8-8.3c-1.6-2-4.2-3.1-6.8-2.9c-2.5.1-4.8 1.5-6.2 3.7c-1.1 1.9-1 4.7-1 7.2c0 2.9.4 8.1 1.6 10.8c.8 1.8 2.9 1.4 3.7 2.6c.7 1.2.8 2.6.3 3.8c-.7 1.6-3.4 2.6-4.3 3.7s-1.1 1.1-1.1 1.1"/>`),
		g.Group(children),
	)
}
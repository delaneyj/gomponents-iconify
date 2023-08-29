package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PnbOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M26.487 26.247v-8.813l5.839 8.813v-8.813m1.767 4.406h2.873m1.534 4.407h-4.407v-8.813H38.5"/><circle cx="19.037" cy="23.202" r="5.768"/><path d="M32.326 34.566V28h-6.241a8.525 8.525 0 0 1-15.573-4.798V20.05c0-2.162-1.012-2.615-1.012-2.615"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/>`),
		g.Group(children),
	)
}
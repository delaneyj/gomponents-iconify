package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redfin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.211 28.329v-6.474l4.289 6.474v-6.474m-6.475 0v6.474m-5.346 0v-6.474h3.237m-3.237 3.237h2.104m-8.338 3.237v-6.474H22.9c1.537 0 2.832 1.295 2.832 2.832v.81c0 1.537-1.295 2.832-2.832 2.832h-1.456Zm-2.026 0h-3.236v-6.474h3.236m-3.236 3.237h2.104M8.5 28.329v-8.357h2.716c1.567 0 2.82 1.254 2.82 2.82s-1.253 2.821-2.82 2.821H8.5m2.818-.002l2.614 2.614"/>`),
		g.Group(children),
	)
}
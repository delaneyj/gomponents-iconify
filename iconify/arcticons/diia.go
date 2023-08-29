package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.632 29V19h-3.274a3.359 3.359 0 0 0 0 6.717h3.274m-3.273 0l-3.274 3.28"/><circle cx="26.288" cy="19.312" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.548 28.006s-1.183 1.286-1.92 1.286h0a.993.993 0 0 1-.923-1.309l1.25-6.314H24.08M20.692 31.5v-4.104H10.357V31.5m8.693-4.105V16.5h-6.38c0 3.172.597 7.873-1.343 10.896"/>`),
		g.Group(children),
	)
}
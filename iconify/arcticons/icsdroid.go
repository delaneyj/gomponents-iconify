package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icsdroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.4 5.5c-1.1 0-1.9.9-1.9 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2V7.4c0-1.1-.9-1.9-2-1.9H7.4zM31.7 9l6 4.5l-6 4.5m1 20h5.2m-5.2-9l2.6-1.4m0 0V38m-8-5.3c1.4 0 2.6-1.2 2.6-2.6h0c0-1.4-1.2-2.6-2.6-2.6h0m0 10.4c1.4 0 2.6-1.2 2.6-2.6h0c0-1.4-1.2-2.6-2.6-2.6h0M23 37c.9.7 2.1 1 3.2.9h1.1M23 28.4c.9-.7 2.1-1 3.2-.9h1.1m-2.7 5.2h2.7m-3.9-19.2h14.3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.2 25.2l-6-4.6l6-4.5m8.3 4.5H10.2"/>`),
		g.Group(children),
	)
}
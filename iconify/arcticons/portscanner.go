package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Portscanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.373 32.5l3.028-17h5.565a4.645 4.645 0 0 1 4.68 5.71a7.115 7.115 0 0 1-6.714 5.709h-5.565m13.307 3.718c.8 1.357 2.018 1.863 3.836 1.863h2.517a5.297 5.297 0 0 0 4.998-4.25h0A3.458 3.458 0 0 0 32.54 24h-2.78a3.458 3.458 0 0 1-3.484-4.25h0a5.297 5.297 0 0 1 4.998-4.25h2.517c1.818 0 3.036.505 3.836 1.863"/>`),
		g.Group(children),
	)
}
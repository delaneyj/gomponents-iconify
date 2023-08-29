package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 5v12.75A3.25 3.25 0 0 0 6.25 21h12a3.25 3.25 0 0 0 3.25-3.25v-8.5A3.251 3.251 0 0 0 19 6.087V5.25A2.25 2.25 0 0 0 16.75 3H5.25a2.25 2.25 0 0 0-2.236 2H3Zm2.25 1a.75.75 0 0 1 0-1.5h11.5a.75.75 0 0 1 .75.75V6H5.25Zm11 7h2a.75.75 0 0 1 0 1.5h-2a.75.75 0 0 1 0-1.5Z"/>`),
		g.Group(children),
	)
}
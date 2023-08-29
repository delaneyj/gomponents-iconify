package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.4 21c-.17.003-.34-.026-.5-.086l8.054-8.057l2.666 2.669l-9.255 5.2A1.998 1.998 0 0 1 5.4 21Zm-1.164-.665a1.9 1.9 0 0 1-.236-.97V4.66a2.13 2.13 0 0 1 .1-.658l8.233 8.235l-8.1 8.1l.003-.002Zm12.179-5.258l-2.841-2.839l3.133-3.132l2.783 1.563c.534.24.892.755.928 1.339a1.574 1.574 0 0 1-.929 1.34l-3.074 1.729Zm-3.461-3.463l-8.34-8.339c.229-.17.506-.26.791-.261c.336.012.664.107.955.277l9.551 5.368l-2.956 2.955h-.001Z"/>`),
		g.Group(children),
	)
}
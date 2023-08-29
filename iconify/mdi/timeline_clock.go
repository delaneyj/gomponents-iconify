package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimelineClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v6H2V2h2M2 22v-6h2v6H2m3-10c0 1.11-.89 2-2 2a2 2 0 1 1 2-2m11-8c4.42 0 8 3.58 8 8s-3.58 8-8 8c-3.6 0-6.64-2.38-7.65-5.65L6 12l2.35-2.35C9.36 6.38 12.4 4 16 4m-1 9l4.53 2.79l.8-1.29l-3.83-2.3V7H15v6Z"/>`),
		g.Group(children),
	)
}
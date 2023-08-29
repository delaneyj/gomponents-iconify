package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M25 38c-7.2 0-13-5.8-13-13c0-3.2 1.2-6.2 3.3-8.6l1.5 1.3C15 19.7 14 22.3 14 25c0 6.1 4.9 11 11 11c1.6 0 3.1-.3 4.6-1l.8 1.8c-1.7.8-3.5 1.2-5.4 1.2zm9.7-4.3l-1.5-1.3c1.8-2 2.8-4.6 2.8-7.3c0-6.1-4.9-11-11-11c-1.6 0-3.1.3-4.6 1l-.8-1.8c1.7-.8 3.5-1.2 5.4-1.2c7.2 0 13 5.8 13 13c0 3.1-1.2 6.2-3.3 8.6z"/><path fill="currentColor" d="M18 24h-2v-6h-6v-2h8zm22 10h-8v-8h2v6h6z"/>`),
		g.Group(children),
	)
}
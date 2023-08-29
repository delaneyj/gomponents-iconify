package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Battery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M48 6h-6V4c0-1.1-.9-2-2-2H24c-1.1 0-2 .9-2 2v2h-6a4.01 4.01 0 0 0-4 4v48c0 2.199 1.799 4 4 4h32c2.199 0 4-1.801 4-4V10c0-2.201-1.801-4-4-4m1 48c0 .541-.459 1-1 1H16c-.543 0-1-.459-1-1V14c0-.543.457-1 1-1h32c.541 0 1 .457 1 1v40"/><path fill="currentColor" d="M44 17H20c-1.1 0-2 .9-2 2v4c0 1.1.9 2 2 2h24c1.1 0 2-.9 2-2v-4c0-1.1-.9-2-2-2m0 13H20c-1.1 0-2 .9-2 2v4c0 1.1.9 2 2 2h24c1.1 0 2-.9 2-2v-4c0-1.1-.9-2-2-2m0 13H20c-1.1 0-2 .9-2 2v4c0 1.1.9 2 2 2h24c1.1 0 2-.9 2-2v-4c0-1.1-.9-2-2-2"/>`),
		g.Group(children),
	)
}
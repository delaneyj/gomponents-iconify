package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskApproved(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 20a6 6 0 1 0-10 4.46V32l4-1.894L28 32v-7.54A5.98 5.98 0 0 0 30 20Zm-4 8.84l-2-.947l-2 .947v-3.19a5.888 5.888 0 0 0 4 0ZM24 24a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Z"/><path fill="currentColor" d="M25 5h-3V4a2.006 2.006 0 0 0-2-2h-8a2.006 2.006 0 0 0-2 2v1H7a2.006 2.006 0 0 0-2 2v21a2.006 2.006 0 0 0 2 2h9v-2H7V7h3v3h12V7h3v5h2V7a2.006 2.006 0 0 0-2-2Zm-5 3h-8V4h8Z"/>`),
		g.Group(children),
	)
}
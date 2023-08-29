package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jiojoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="30" height="22" x="3.57" y="10.62" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.57 25.541l7.92 4.573a1 1 0 0 0 1.5-.866V15.392a1 1 0 0 0-1.5-.866l-7.92 4.573m-15-4.479a3.394 3.394 0 1 1-3.394 3.394a3.394 3.394 0 0 1 3.394-3.394Zm0 8.44c3.772 0 6.789 1.057 6.789 2.316v3.244H11.78v-3.229c0-1.274 3.018-2.331 6.789-2.331Z"/>`),
		g.Group(children),
	)
}
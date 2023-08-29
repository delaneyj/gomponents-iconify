package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TataOneMg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2ZM11 23.67l2.877-1.566m0 0v11.507M11 10.505h6.549m-3.274 7.715v-7.715"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37 25.988v8.63a2.877 2.877 0 0 1-2.877 2.877h0a2.868 2.868 0 0 1-2.034-.843"/><rect width="5.754" height="7.624" x="31.246" y="25.988" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.877" transform="rotate(180 34.123 29.8)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.862 28.864a2.877 2.877 0 0 1 2.877-2.877h0a2.877 2.877 0 0 1 2.877 2.877v4.747m-5.754-7.623v7.623m5.754-4.747a2.877 2.877 0 0 1 2.877-2.877h0a2.877 2.877 0 0 1 2.877 2.877v4.747M17.549 18.197l3.212-7.692m3.214 7.715l-3.214-7.715m3.264 0h6.549m-3.275 7.715v-7.715m3.275 7.692l3.211-7.692M37 18.22l-3.215-7.715M16.681 22.104H37m-8.631 15.391H13.877"/>`),
		g.Group(children),
	)
}
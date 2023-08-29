package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.78 16.035a6.387 6.387 0 1 0-7.559 0m7.559 0l4.53 2.983a8.461 8.461 0 0 1-.768 2.618H16.458a8.461 8.461 0 0 1-.768-2.618l4.53-2.983m7.309 5.601c0 9.073 9.388 10.4 9.388 19.194c-1.745 2.19-7.93 2.67-12.917 2.67s-11.172-.48-12.917-2.67c0-8.794 9.388-10.12 9.388-19.194"/>`),
		g.Group(children),
	)
}
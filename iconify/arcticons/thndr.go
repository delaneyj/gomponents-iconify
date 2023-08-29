package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thndr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.305 8.132L4.52 39.867a.297.297 0 0 0 .173.381a.293.293 0 0 0 .104.02H8.76a.612.612 0 0 0 .574-.4L21.077 8.133a.297.297 0 0 0-.277-.4l-3.92-.002a.612.612 0 0 0-.574.4Zm3.507 3.42h12.885s.702.288.532.923c-.318 1.234-1.122 3.032-1.335 3.865c.255.137 11.606 0 11.606 0l-20.513 9.593l1.63-4.243s-.13-.513-1.097-.513h-7.152"/>`),
		g.Group(children),
	)
}
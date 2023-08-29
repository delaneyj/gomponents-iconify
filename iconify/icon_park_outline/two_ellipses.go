package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoEllipses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40.579 7.349c3.765 3.764-.622 14.255-9.799 23.431c-9.176 9.177-19.667 13.564-23.431 9.799c-3.765-3.765.622-14.255 9.798-23.432C26.324 7.971 36.814 3.584 40.58 7.35Z"/><path d="M7.485 7.349c-3.764 3.764.623 14.255 9.799 23.431c9.176 9.177 19.667 13.564 23.432 9.799c3.764-3.765-.623-14.255-9.799-23.432C21.741 7.971 11.25 3.584 7.485 7.35Z"/></g>`),
		g.Group(children),
	)
}
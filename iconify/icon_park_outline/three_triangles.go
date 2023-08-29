package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeTriangles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="m17 7l7 12.124L29.99 29.5H4.01L17 7Z"/><path d="m31 7l12.99 22.5H18.01L24 19.124L31 7ZM11.01 41.5h25.98l-7-12H18.01l-7 12Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M29.99 29.5L24 19.124M18.01 29.5l-7 12h25.98l-7-12H18.01Zm0 0h25.98L31 7l-7 12.124L18.01 29.5Zm0 0L24 19.124L18.01 29.5Zm0 0h11.98h-11.98Zm11.98 0H4.01L17 7l7 12.124L29.99 29.5Z"/></g>`),
		g.Group(children),
	)
}
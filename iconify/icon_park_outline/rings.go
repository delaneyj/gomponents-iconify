package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path stroke-linecap="round" d="M13 43c4.42 0 8-3.58 8-8s-3.58-8-8-8s-8 3.58-8 8s3.58 8 8 8Zm22 0c4.42 0 8-3.58 8-8s-3.58-8-8-8s-8 3.58-8 8s3.58 8 8 8ZM6 5h36"/><path stroke-linecap="square" d="M13 27V5m22 22V5"/><path stroke-linecap="round" d="M9 19h8m14 0h8"/></g>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 6h36v36H6V6Z"/><path d="M36 24c0 6.627-5.373 12-12 12s-12-5.373-12-12s5.373-12 12-12s12 5.373 12 12ZM4 4l40 40m0-40L4 44"/></g>`),
		g.Group(children),
	)
}
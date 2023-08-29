package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alpinedotjs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m24 12l-5.72 5.746l-5.724-5.741l5.724-5.75L24 12zM5.72 6.254L0 12l5.72 5.746h11.44L5.72 6.254z"/>`),
		g.Group(children),
	)
}
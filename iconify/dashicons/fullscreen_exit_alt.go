package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullscreenExitAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.4 2L2 3.4l2.8 2.8L3 8h5V3L6.2 4.8L3.4 2zm11.8 4.2L18 3.4L16.6 2l-2.8 2.8L12 3v5h5l-1.8-1.8zM4.8 13.8L2 16.6L3.4 18l2.8-2.8L8 17v-5H3l1.8 1.8zM17 12h-5v5l1.8-1.8l2.8 2.8l1.4-1.4l-2.8-2.8L17 12z"/>`),
		g.Group(children),
	)
}
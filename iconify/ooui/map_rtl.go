package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 1v16l6 2l6-2l6 2V3l-6-2l-6 2zm12 2l4 1.36v11.9l-4-1.36zM3 3.74L7 5.1V17l-4-1.36z"/>`),
		g.Group(children),
	)
}
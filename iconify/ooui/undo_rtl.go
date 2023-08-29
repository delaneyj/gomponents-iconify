package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UndoRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 3v4h-1C6 7 1 10 1 16v1h3v-1c0-4 3-6 7-6h1v4l7-5.5z"/>`),
		g.Group(children),
	)
}
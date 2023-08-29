package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Search(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.2 13.6a7 7 0 1 1 1.4-1.4l5.4 5.4l-1.4 1.4zM3 8a5 5 0 1 0 10 0A5 5 0 0 0 3 8z"/>`),
		g.Group(children),
	)
}
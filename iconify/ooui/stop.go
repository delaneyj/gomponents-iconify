package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<rect width="16" height="16" x="2" y="2" fill="currentColor" rx="1"/>`),
		g.Group(children),
	)
}
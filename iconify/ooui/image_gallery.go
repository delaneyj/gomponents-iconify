package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageGallery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2zm0 11l3.5-4.5l2.5 3l3.5-4.5l4.5 6zM16 2a2 2 0 0 1 2 2H2a2 2 0 0 1 2-2z"/>`),
		g.Group(children),
	)
}
package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoldA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.5 15h-7L5 19H1L8 1h4l7 18h-4Zm-6-3h5L10 4Z"/>`),
		g.Group(children),
	)
}
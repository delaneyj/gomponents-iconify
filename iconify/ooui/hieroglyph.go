package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hieroglyph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 11h-3.75l2.55-3.4a4.75 4.75 0 1 0-7.6 0L8.75 11H5v2h4v7h2v-7h4zM7.54 3.52A2.75 2.75 0 1 1 12.2 6.4L10 9.33L7.8 6.4a2.69 2.69 0 0 1-.26-2.88z"/>`),
		g.Group(children),
	)
}
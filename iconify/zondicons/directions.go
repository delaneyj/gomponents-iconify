package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Directions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10 0l10 10l-10 10L0 10L10 0zM6 10v3h2v-3h3v3l4-4l-4-4v3H8a2 2 0 0 0-2 2z"/>`),
		g.Group(children),
	)
}
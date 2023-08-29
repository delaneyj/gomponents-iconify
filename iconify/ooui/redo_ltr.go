package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedoLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 8.5L12 3v11zM12 7v3h-1c-4 0-7 2-7 6v1H1v-1c0-6 5-9 10-9z"/>`),
		g.Group(children),
	)
}
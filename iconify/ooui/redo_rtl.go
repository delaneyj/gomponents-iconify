package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedoRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 8.5L8 3v11zM8 7v3h1c4 0 7 2 7 6v1h3v-1c0-6-5-9-10-9z"/>`),
		g.Group(children),
	)
}
package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 1L0 10h3v9h4v-4.6c0-1.47 1.31-2.66 3-2.66s3 1.19 3 2.66V19h4v-9h3L10 1z"/>`),
		g.Group(children),
	)
}
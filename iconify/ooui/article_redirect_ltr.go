package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArticleRedirectLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 1a2 2 0 0 0-2 2v1c0 5 2 8 7 8V9l5 4l-5 4v-3c-3.18 0-5.51-.85-7-2.68V17a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}
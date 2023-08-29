package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArticleRedirectRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 17c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2v-5.7c-1.5 1.8-3.8 2.7-7 2.7v3l-5-4l5-4v3c5 0 7-3 7-8V3c0-1.1-.9-2-2-2H5c-1.1 0-2 .9-2 2z"/>`),
		g.Group(children),
	)
}
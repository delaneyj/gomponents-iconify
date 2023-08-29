package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArticleDisambiguationLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 1H5c-1.1 0-2 .9-2 2v6h4.6l3.7-3.7L10 4h4v4l-1.3-1.3L9.4 10l3.3 3.3L14 12v4h-4l1.3-1.3L7.6 11H3v6c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2z"/>`),
		g.Group(children),
	)
}
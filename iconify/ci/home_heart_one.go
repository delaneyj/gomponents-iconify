package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeHeartOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 22H5a1 1 0 0 1-1-1v-9.643a1.01 1.01 0 0 1 .293-.65l7-7a1 1 0 0 1 1.415 0l7 7a.994.994 0 0 1 .292.707V21a1 1 0 0 1-1 1Zm-7.622-4.525c.22.18.433.356.622.525c.2-.177.424-.36.662-.554l.057-.047c1.166-.951 2.618-2.134 2.618-3.6A1.8 1.8 0 0 0 13.5 12a2.008 2.008 0 0 0-1.5.667A2.007 2.007 0 0 0 10.5 12a1.8 1.8 0 0 0-1.835 1.8c0 1.461 1.44 2.636 2.6 3.58h-.001h.005l.018.015h.005l.008.019l.052.043l.013.01h.006l.007.008Z"/>`),
		g.Group(children),
	)
}
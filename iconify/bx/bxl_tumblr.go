package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxlTumblr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M14.078 20.953c-2.692 0-4.699-1.385-4.699-4.7v-5.308H6.931V8.07c2.694-.699 3.821-3.017 3.95-5.023h2.796v4.558h3.263v3.34h-3.263v4.622c0 1.386.699 1.864 1.813 1.864h1.58v3.522h-2.992z" fill="currentColor"/>`),
		g.Group(children),
	)
}
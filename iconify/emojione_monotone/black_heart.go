package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M61.072 18.16c-6.395-16.919-27.154-9.328-29.074-.879c-2.64-9.004-22.89-15.721-29.07.891C-3.953 36.674 29.598 53.279 31.999 56c2.397-2.162 35.951-19.639 29.073-37.84"/>`),
		g.Group(children),
	)
}
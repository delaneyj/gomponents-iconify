package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M1 0v1c3.32 0 6 2.68 6 6h1c0-3.86-3.14-7-7-7zm0 2v1c2.21 0 4 1.79 4 4h1c0-2.76-2.24-5-5-5zm0 2v1c1.11 0 2 .89 2 2h1c0-1.65-1.35-3-3-3zm0 2c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1z"/>`),
		g.Group(children),
	)
}
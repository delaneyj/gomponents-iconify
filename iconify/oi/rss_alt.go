package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v2c3.33 0 6 2.67 6 6h2c0-4.41-3.59-8-8-8zm0 3v2c1.67 0 3 1.33 3 3h2c0-2.75-2.25-5-5-5zm0 3v2h2a2 2 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}
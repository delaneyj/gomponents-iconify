package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v4c9.93 0 18 8.07 18 18h4C27 14.85 17.15 5 5 5zm0 7v4c6.07 0 11 4.93 11 11h4c0-8.28-6.72-15-15-15zm3 9a3 3 0 0 0 0 6a3 3 0 0 0 0-6z"/>`),
		g.Group(children),
	)
}
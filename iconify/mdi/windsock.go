package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windsock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 5v8l15-2V7L7 5m3 1.91l3 .4v3.38l-3 .4V6.91m6 .8l3 .4v1.78l-3 .4V7.71M5 10v1h1v1H5v9H3V4c0-.55.45-1 1-1s1 .45 1 1v2h1v1H5v3Z"/>`),
		g.Group(children),
	)
}
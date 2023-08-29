package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoinPound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.5 1a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15zm0 13.5a6 6 0 1 1 0-12a6 6 0 0 1 0 12z"/><path fill="currentColor" d="M9.5 11H6V9h1.5a.5.5 0 0 0 0-1H6v-.5a1.502 1.502 0 0 1 2.8-.75a.499.499 0 1 0 .865-.501A2.51 2.51 0 0 0 7.5 4.999a2.503 2.503 0 0 0-2.5 2.5v.5h-.5a.5.5 0 0 0 0 1H5v3h4.5a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}
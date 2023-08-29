package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoinDollar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.5 1a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15zm0 13.5a6 6 0 1 1 0-12a6 6 0 0 1 0 12zM8 8V6h2V5H8V4H7v1H5v4h2v2H5v1h2v1h1v-1h2V8H8zM7 8H6V6h1v2zm2 3H8V9h1v2z"/>`),
		g.Group(children),
	)
}
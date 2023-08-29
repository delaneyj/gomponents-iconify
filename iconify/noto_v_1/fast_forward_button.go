package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="m6.24 98.99l54.93-31.71v31.71L121.76 64L61.17 29.01v31.72L6.24 29.01z"/>`),
		g.Group(children),
	)
}
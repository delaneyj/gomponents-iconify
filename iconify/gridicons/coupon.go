package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coupon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 16v2h-2v-2h2zm3-3h2v-2h-2v2zm2 8h-2v2h2v-2zm3-5v2h2v-2h-2zm-1-3a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2zm1 7a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2zm-7 1a1 1 0 0 1-1-1h-2a3 3 0 0 0 3 3v-2zm3.21-5.21a2 2 0 0 1-2.828.002l-.002-.002L10 11.41l-1.43 1.44c.279.506.427 1.073.43 1.65A3.5 3.5 0 1 1 5.5 11a3.454 3.454 0 0 1 1.65.43L8.59 10L7.15 8.57A3.454 3.454 0 0 1 5.5 9A3.5 3.5 0 1 1 9 5.5a3.454 3.454 0 0 1-.43 1.65L10 8.59l3.88-3.88a2 2 0 0 1 2.828-.002l.002.002l-5.3 5.29l5.8 5.79zM5.5 7a1.5 1.5 0 1 0-.001-3.001A1.5 1.5 0 0 0 5.5 7zM7 14.5a1.5 1.5 0 1 0-3.001.001A1.5 1.5 0 0 0 7 14.5z"/>`),
		g.Group(children),
	)
}
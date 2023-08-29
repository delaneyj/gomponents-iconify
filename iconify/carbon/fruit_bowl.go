package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FruitBowl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 15a6.003 6.003 0 0 0-3.107-5.253A3.98 3.98 0 0 0 24 7h-2a2.002 2.002 0 0 1-2 2a6.004 6.004 0 0 0-5.995 5.892A7 7 0 0 1 12 10a3.996 3.996 0 0 0-3-3.858V4H7v2.142A3.996 3.996 0 0 0 4 10v5H2v1a14 14 0 0 0 28 0v-1Zm-6-4a4.005 4.005 0 0 1 4 4h-8a4.005 4.005 0 0 1 4-4ZM6 10a2 2 0 1 1 4 0a8.991 8.991 0 0 0 1.532 5H6Zm10 18A12.017 12.017 0 0 1 4.041 17H27.96A12.017 12.017 0 0 1 16 28Z"/>`),
		g.Group(children),
	)
}
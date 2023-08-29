package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagPercentTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 5v2H5.5A1.5 1.5 0 0 0 4 8.5V18a4 4 0 0 0 4 4h4.168a1.75 1.75 0 0 1 .345-1.987l.708-.71A3.998 3.998 0 0 1 13 18v-.337a3.5 3.5 0 0 1 0-6.326V5c0-.334-.055-.656-.156-.956A1.5 1.5 0 0 1 15.5 5v2h-1v4a3.5 3.5 0 0 1 3.5 3.5v.025l2-2V8.5A1.5 1.5 0 0 0 18.5 7H17V5a3 3 0 0 0-5-2.236A3 3 0 0 0 7 5Zm1.5 2V5a1.5 1.5 0 1 1 3 0v2h-3Zm6 10a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-3.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm6 9.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-3.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm1.28-5.22a.75.75 0 1 0-1.06-1.06l-7.5 7.5a.75.75 0 1 0 1.06 1.06l7.5-7.5Z"/>`),
		g.Group(children),
	)
}
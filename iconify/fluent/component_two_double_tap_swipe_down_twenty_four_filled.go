package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentTwoDoubleTapSwipeDownTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 15.71v-1.585A5.502 5.502 0 0 1 12 3.5a5.5 5.5 0 0 1 2.21 10.538l-.21.087l.001 1.585a7 7 0 1 0-4.247-.078l.246.078Zm0-2.678v-1.796a3 3 0 1 1 4.138-.132l-.137.131v1.797a4.5 4.5 0 1 0-4.192-.1l.19.1Zm1.387 8.758a1 1 0 0 0 1.226 0l.094-.083l3-3l.083-.094a1 1 0 0 0 0-1.226l-.083-.094l-.094-.083a1 1 0 0 0-1.226 0l-.094.083L13 18.585V9a1 1 0 0 0-1.993-.117L11 9v9.585l-1.293-1.292l-.094-.083a1 1 0 0 0-1.403 1.403l.083.094l3 3l.094.083Z"/>`),
		g.Group(children),
	)
}
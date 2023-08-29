package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckCircleFillSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 16A8 8 0 1 1 8 0a8 8 0 0 1 0 16Zm3.78-9.72a.751.751 0 0 0-.018-1.042a.751.751 0 0 0-1.042-.018L6.75 9.19L5.28 7.72a.751.751 0 0 0-1.042.018a.751.751 0 0 0-.018 1.042l2 2a.75.75 0 0 0 1.06 0Z"/>`),
		g.Group(children),
	)
}
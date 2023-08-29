package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoryTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 5a8.96 8.96 0 0 0-5.658 2H10a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1.013V4a1 1 0 0 1 2 0v1.514A10.957 10.957 0 0 1 14 3c6.075 0 11 4.925 11 11s-4.925 11-11 11S3 20.075 3 14c0-.37.018-.737.054-1.099a1 1 0 1 1 1.99.198A9 9 0 1 0 14 5Zm1 3a1 1 0 1 0-2 0v6a1 1 0 0 0 1 1h2.5a1 1 0 1 0 0-2H15V8Z"/>`),
		g.Group(children),
	)
}
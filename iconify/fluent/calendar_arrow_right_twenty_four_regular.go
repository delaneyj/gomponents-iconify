package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarArrowRightTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 6a3 3 0 0 0-3-3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h6.022a6.471 6.471 0 0 1-.709-1.5H6A1.5 1.5 0 0 1 4.5 18V8.5h15v2.813a6.471 6.471 0 0 1 1.5.709V6ZM6 4.5h12A1.5 1.5 0 0 1 19.5 6v1h-15V6A1.5 1.5 0 0 1 6 4.5Zm17 13a5.5 5.5 0 1 0-11 0a5.5 5.5 0 0 0 11 0Zm-5.354-2.146a.5.5 0 0 1 .708-.708l2.5 2.5a.5.5 0 0 1 0 .708l-2.5 2.5a.5.5 0 0 1-.708-.708L19.293 18H14.5a.5.5 0 0 1 0-1h4.793l-1.647-1.646Z"/>`),
		g.Group(children),
	)
}
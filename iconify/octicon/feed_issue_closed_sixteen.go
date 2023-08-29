package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeedIssueClosedSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8.5 0a8 8 0 1 1 0 16a8 8 0 0 1 0-16Zm3.457 6.957a.999.999 0 1 0-1.414-1.414L7.75 8.336L6.457 7.043a.999.999 0 1 0-1.414 1.414l2 2a.999.999 0 0 0 1.414 0Z"/>`),
		g.Group(children),
	)
}
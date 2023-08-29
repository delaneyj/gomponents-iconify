package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.043 4.076c-.275-1.03.783-1.91 1.746-1.451l21.497 10.249c.95.452.95 1.804 0 2.256L3.79 25.38c-.963.458-2.021-.422-1.746-1.452L4.7 14.002L2.043 4.076Zm4.008 10.676l-2.418 9.04l20.535-9.79l-20.535-9.79l2.418 9.04H17.25a.75.75 0 0 1 0 1.5H6.05Z"/>`),
		g.Group(children),
	)
}
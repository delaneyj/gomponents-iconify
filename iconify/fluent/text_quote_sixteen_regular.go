package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextQuoteSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm.885 2.794c-.23 1.592-.852 2.966-2.239 4.352a.5.5 0 0 0 .708.708C12.473 9.734 13 7.592 13 5a2 2 0 1 0-1.115 1.794ZM5 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm.885 2.794c-.23 1.592-.852 2.966-2.239 4.352a.5.5 0 0 0 .708.708C6.473 9.734 7 7.592 7 5a2 2 0 1 0-1.115 1.794Z"/>`),
		g.Group(children),
	)
}
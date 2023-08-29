package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateThreeHundredSixty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6C6.3 6 2 8.15 2 11c0 2.45 3.19 4.38 7.71 4.88l-.42.41a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2-2a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76a1 1 0 0 0-.21-.33l-2-2a1 1 0 0 0-1.42 1.42l.12.11C6 13.34 4 12 4 11c0-1.22 3.12-3 8-3s8 1.78 8 3c0 .83-1.45 2-4.21 2.6a1 1 0 0 0-.79 1.19a1 1 0 0 0 1.19.77c3.65-.8 5.81-2.5 5.81-4.56c0-2.85-4.3-5-10-5Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleOffSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m11.874 12.582l2.272 2.272a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708l2.272 2.272a6 6 0 0 0 8.456 8.456Zm-.71-.71a5 5 0 0 1-7.036-7.036l7.036 7.035ZM13 8c0 .834-.204 1.621-.566 2.313l.735.735A6 6 0 0 0 4.952 2.83l.735.735A5 5 0 0 1 13 8Z"/>`),
		g.Group(children),
	)
}
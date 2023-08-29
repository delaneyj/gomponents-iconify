package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 3H7.4S3 2.8 3 7.3C3 10.8 5 14 5 14s-.4-7 2.3-7H11v3l5-5l-5-5v3z"/><path fill="currentColor" d="M14 9v6H1V2h9V1H0v15h15V8z"/>`),
		g.Group(children),
	)
}
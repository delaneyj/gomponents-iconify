package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LatinCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#f79229" d="M69 2H58c-1.1 0-2 .9-2 2v30H29c-1.1 0-2 .9-2 2v11c0 1.1.9 2 2 2h27v75c0 1.1.9 2 2 2h11c1.1 0 2-.9 2-2V49h28c1.1 0 2-.9 2-2V36c0-1.1-.9-2-2-2H71V4c0-1.1-.9-2-2-2z"/>`),
		g.Group(children),
	)
}
package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteLargeSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#e0e0e0" d="M6.93 1C3.67 1 1 3.67 1 6.93v114.14c0 3.26 2.67 5.93 5.93 5.93h114.14c3.26 0 5.93-2.67 5.93-5.93V6.93c0-3.26-2.67-5.93-5.93-5.93H6.93z"/>`),
		g.Group(children),
	)
}
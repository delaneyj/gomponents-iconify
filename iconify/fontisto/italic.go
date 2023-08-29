package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.182 0H5.423a1.091 1.091 0 0 0 0 2.182h.032h-.002h3.033L4.559 21.819h-3.5a1.091 1.091 0 0 0 0 2.182h.032h-.002h8.727a1.092 1.092 0 0 0 .002-2.182H6.783L10.71 2.182h3.469A1.092 1.092 0 0 0 14.181 0h-.002z"/>`),
		g.Group(children),
	)
}
package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="m64 22.95l31.13 42.47H75.71v39.63H52.48V65.42H32.86L64 22.95z"/>`),
		g.Group(children),
	)
}
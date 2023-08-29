package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="M98.25 98.18L44.51 89.9l14.18-14.18l-28.94-28.93l16.96-16.96l28.94 28.93l14.32-14.32l8.28 53.74z"/>`),
		g.Group(children),
	)
}
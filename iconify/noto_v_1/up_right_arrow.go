package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="M98.18 29.76L89.9 83.49L75.72 69.31L46.78 98.24L29.82 81.28l28.94-28.93l-14.32-14.33l53.74-8.26z"/>`),
		g.Group(children),
	)
}
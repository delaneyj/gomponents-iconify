package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrowCurvingLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="M40.37 81.55V66.3L7.34 90.51l33.03 24.21V99.55h18.99c37.33 0 61.3-10.93 61.3-43.14s-23.97-43.14-61.3-43.14H16.6v17.94h42.76c31.8 0 43.36 6.83 43.36 25.2c0 18.38-11.55 25.21-43.36 25.21H40.37v-.07z"/>`),
		g.Group(children),
	)
}
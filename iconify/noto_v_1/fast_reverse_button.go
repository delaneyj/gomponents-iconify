package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastReverseButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="M121.76 29.01L66.84 60.72V29.01L6.24 64l60.6 34.99V67.28l54.92 31.71z"/>`),
		g.Group(children),
	)
}
package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastDownButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="m29.01 6.24l31.71 54.92H29.01L64 121.76l34.99-60.6H67.28L98.99 6.24z"/>`),
		g.Group(children),
	)
}
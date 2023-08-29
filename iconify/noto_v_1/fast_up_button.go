package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastUpButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="M98.99 121.76L67.28 66.84h31.71L64 6.24l-34.99 60.6h31.71l-31.71 54.92z"/>`),
		g.Group(children),
	)
}
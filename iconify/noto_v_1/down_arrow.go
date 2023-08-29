package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="M64 105.05L32.86 62.58h19.42V22.95h23.23v39.63h19.62L64 105.05z"/>`),
		g.Group(children),
	)
}
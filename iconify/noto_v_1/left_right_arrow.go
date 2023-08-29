package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="M44.3 51.91V31.85L.46 63.99L44.3 96.15V76.1h39.4v20.05L127.54 64L83.7 31.85v20.06z"/>`),
		g.Group(children),
	)
}
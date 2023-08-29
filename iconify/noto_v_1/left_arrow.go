package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="m22.95 64l42.46-31.14v19.42h39.64v23.23H65.41v19.61L22.95 64z"/>`),
		g.Group(children),
	)
}
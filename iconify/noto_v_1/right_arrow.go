package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="M108.58 64L62.47 97.81V76.72H19.42V51.49h43.04v-21.3L108.58 64z"/>`),
		g.Group(children),
	)
}
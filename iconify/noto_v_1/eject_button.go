package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="M28.08 89.99h72v18h-72v-18zm-.16-12.11h72L63.92 20l-36 57.88z"/>`),
		g.Group(children),
	)
}
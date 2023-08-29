package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAlphaDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m206.392 382.863l-51.883 51.882V17.177h-32v417.568l-51.882-51.882L48 405.49L138.509 496l90.51-90.51l-22.627-22.627zm84.921 74.314h144v-32H326.274l109.039-100.732v-43.268h-136v32h101.04l-109.04 100.732v43.268zm52.468-408l-58.666 176h33.73l18.667-56h59.6l18.666 56h33.731l-58.666-176Zm4.4 88l18.667-56h.935l18.667 56Z"/>`),
		g.Group(children),
	)
}
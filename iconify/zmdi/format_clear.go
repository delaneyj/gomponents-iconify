package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatClear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m27 43l6 5l308 309l-27 27l-121-121l-33 78H96l53-123L0 70zm58 0h299v64H260l-34 80l-45-44l16-36h-52L85 47v-4z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlidersBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M68 102.06V40a12 12 0 0 0-24 0v62.06a36 36 0 0 0 0 67.88V216a12 12 0 0 0 24 0v-46.06a36 36 0 0 0 0-67.88ZM56 148a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm108-60a36.07 36.07 0 0 0-24-33.94V40a12 12 0 0 0-24 0v14.06a36 36 0 0 0 0 67.88V216a12 12 0 0 0 24 0v-94.06A36.07 36.07 0 0 0 164 88Zm-36 12a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm108 68a36.07 36.07 0 0 0-24-33.94V40a12 12 0 0 0-24 0v94.06a36 36 0 0 0 0 67.88V216a12 12 0 0 0 24 0v-14.06A36.07 36.07 0 0 0 236 168Zm-36 12a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}
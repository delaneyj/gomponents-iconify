package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutLineVerticalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M88.61 51.06a8 8 0 0 1 1.73-8.72l32-32a8 8 0 0 1 11.32 0l32 32A8 8 0 0 1 160 56h-24v40a8 8 0 0 1-16 0V56H96a8 8 0 0 1-7.39-4.94ZM216 120H40a8 8 0 0 0 0 16h176a8 8 0 0 0 0-16Zm-56 80h-24v-40a8 8 0 0 0-16 0v40H96a8 8 0 0 0-5.66 13.66l32 32a8 8 0 0 0 11.32 0l32-32A8 8 0 0 0 160 200Z"/>`),
		g.Group(children),
	)
}
package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraduationCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m85 217l150 82l149-82v86l-149 81l-150-81v-86zM235 0l234 128v171h-42V151L235 256L0 128z"/>`),
		g.Group(children),
	)
}
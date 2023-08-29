package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M412 319v100H312l-71 70l-70-70H71V319L0 248l71-71V77h100l70-70l71 70h100v100l71 71zm-171 57q53 0 90.5-37.5T369 248t-37.5-90.5T241 120t-90.5 37.5T113 248t37.5 90.5T241 376z"/>`),
		g.Group(children),
	)
}
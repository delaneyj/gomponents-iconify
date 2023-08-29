package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m412 177l71 71l-71 71v100H312l-71 70l-70-70H71V319L0 248l71-71V77h100l70-70l71 70h100v100zM241 376q53 0 90.5-37.5T369 248t-37.5-90.5T241 120q-28 0-53 12q33 15 54 46.5t21 69.5t-21 69.5t-54 46.5q25 12 53 12z"/>`),
		g.Group(children),
	)
}
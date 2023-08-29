package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProductHunt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1150 762q0 56-39.5 95t-95.5 39H762V627h253q56 0 95.5 39.5T1150 762zm179 0q0-130-91.5-222T1015 448H582v896h180v-269h253q130 0 222-91.5t92-221.5zm463 134q0 182-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348z"/>`),
		g.Group(children),
	)
}
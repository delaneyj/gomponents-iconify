package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoundingBoxLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 94a14 14 0 0 0 14-14V48a14 14 0 0 0-14-14h-32a14 14 0 0 0-14 14v10H94V48a14 14 0 0 0-14-14H48a14 14 0 0 0-14 14v32a14 14 0 0 0 14 14h10v68H48a14 14 0 0 0-14 14v32a14 14 0 0 0 14 14h32a14 14 0 0 0 14-14v-10h68v10a14 14 0 0 0 14 14h32a14 14 0 0 0 14-14v-32a14 14 0 0 0-14-14h-10V94Zm-34-46a2 2 0 0 1 2-2h32a2 2 0 0 1 2 2v32a2 2 0 0 1-2 2h-32a2 2 0 0 1-2-2ZM46 80V48a2 2 0 0 1 2-2h32a2 2 0 0 1 2 2v32a2 2 0 0 1-2 2H48a2 2 0 0 1-2-2Zm36 128a2 2 0 0 1-2 2H48a2 2 0 0 1-2-2v-32a2 2 0 0 1 2-2h32a2 2 0 0 1 2 2Zm128-32v32a2 2 0 0 1-2 2h-32a2 2 0 0 1-2-2v-32a2 2 0 0 1 2-2h32a2 2 0 0 1 2 2Zm-24-14h-10a14 14 0 0 0-14 14v10H94v-10a14 14 0 0 0-14-14H70V94h10a14 14 0 0 0 14-14V70h68v10a14 14 0 0 0 14 14h10Z"/>`),
		g.Group(children),
	)
}
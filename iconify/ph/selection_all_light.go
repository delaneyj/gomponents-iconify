package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectionAllLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M106 40a6 6 0 0 1 6-6h32a6 6 0 0 1 0 12h-32a6 6 0 0 1-6-6Zm38 170h-32a6 6 0 0 0 0 12h32a6 6 0 0 0 0-12Zm64-176h-24a6 6 0 0 0 0 12h24a2 2 0 0 1 2 2v24a6 6 0 0 0 12 0V48a14 14 0 0 0-14-14Zm8 72a6 6 0 0 0-6 6v32a6 6 0 0 0 12 0v-32a6 6 0 0 0-6-6Zm0 72a6 6 0 0 0-6 6v24a2 2 0 0 1-2 2h-24a6 6 0 0 0 0 12h24a14 14 0 0 0 14-14v-24a6 6 0 0 0-6-6ZM40 150a6 6 0 0 0 6-6v-32a6 6 0 0 0-12 0v32a6 6 0 0 0 6 6Zm32 60H48a2 2 0 0 1-2-2v-24a6 6 0 0 0-12 0v24a14 14 0 0 0 14 14h24a6 6 0 0 0 0-12ZM40 78a6 6 0 0 0 6-6V48a2 2 0 0 1 2-2h24a6 6 0 0 0 0-12H48a14 14 0 0 0-14 14v24a6 6 0 0 0 6 6Zm136 104H80a6 6 0 0 1-6-6V80a6 6 0 0 1 6-6h96a6 6 0 0 1 6 6v96a6 6 0 0 1-6 6Zm-6-96H86v84h84Z"/>`),
		g.Group(children),
	)
}
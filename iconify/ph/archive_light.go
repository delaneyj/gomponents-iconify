package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 50H32a14 14 0 0 0-14 14v24a14 14 0 0 0 14 14h2v90a14 14 0 0 0 14 14h160a14 14 0 0 0 14-14v-90h2a14 14 0 0 0 14-14V64a14 14 0 0 0-14-14Zm-14 142a2 2 0 0 1-2 2H48a2 2 0 0 1-2-2v-90h164Zm16-104a2 2 0 0 1-2 2H32a2 2 0 0 1-2-2V64a2 2 0 0 1 2-2h192a2 2 0 0 1 2 2ZM98 136a6 6 0 0 1 6-6h48a6 6 0 0 1 0 12h-48a6 6 0 0 1-6-6Z"/>`),
		g.Group(children),
	)
}
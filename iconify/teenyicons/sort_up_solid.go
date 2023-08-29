package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortUpSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.146.146a.5.5 0 0 1 .708 0l3 3l-.708.708L4 1.707V15H3V1.707L.854 3.854l-.708-.708l3-3ZM15 4H9V3h6v1Zm-2 4H9V7h4v1Zm-2 4H9v-1h2v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
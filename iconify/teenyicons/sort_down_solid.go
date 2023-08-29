package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortDownSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 13.293V0h1v13.293l2.146-2.147l.708.708l-3 3a.5.5 0 0 1-.708 0l-3-3l.708-.708L3 13.293ZM15 4H9V3h6v1Zm-2 4H9V7h4v1Zm-2 4H9v-1h2v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
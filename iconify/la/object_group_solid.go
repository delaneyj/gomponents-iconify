package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectGroupSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v4h1v14H5v4h4v-1h14v1h4v-4h-1V9h1V5h-4v1H9V5zm4 3h14v1h1v14h-1v1H9v-1H8V9h1zm1 2v8h4v4h8v-8h-4v-4zm2 2h4v4h-4zm6 4h2v4h-4v-2h2z"/>`),
		g.Group(children),
	)
}
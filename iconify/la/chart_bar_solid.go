package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 4v24h8V4zm2 2h4v20h-4zM3 10v18h8V10zm2 2h4v14H5zm7 4v12h8V16zm2 2h4v8h-4z"/>`),
		g.Group(children),
	)
}
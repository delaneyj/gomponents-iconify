package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataSet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 13V4h-2v2h-3v2h3v5h-3v2h8v-2h-3zM8.5 6A3.5 3.5 0 1 1 5 9.5A3.5 3.5 0 0 1 8.5 6m0-2A5.5 5.5 0 1 0 14 9.5A5.5 5.5 0 0 0 8.5 4zm0 16A3.5 3.5 0 1 1 5 23.5A3.5 3.5 0 0 1 8.5 20m0-2a5.5 5.5 0 1 0 5.5 5.5A5.5 5.5 0 0 0 8.5 18zm15 2a3.5 3.5 0 1 1-3.5 3.5a3.5 3.5 0 0 1 3.5-3.5m0-2a5.5 5.5 0 1 0 5.5 5.5a5.5 5.5 0 0 0-5.5-5.5z"/>`),
		g.Group(children),
	)
}
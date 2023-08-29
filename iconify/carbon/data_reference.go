package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 13V4h-2v2h-3v2h3v5h-3v2h8v-2h-3zM8.5 6A3.5 3.5 0 1 1 5 9.5A3.504 3.504 0 0 1 8.5 6m0-2A5.5 5.5 0 1 0 14 9.5A5.5 5.5 0 0 0 8.5 4zm15 16a3.5 3.5 0 1 1-3.5 3.5a3.504 3.504 0 0 1 3.5-3.5m0-2a5.5 5.5 0 1 0 5.5 5.5a5.5 5.5 0 0 0-5.5-5.5zM6 19v2h3.586L4 26.586L5.414 28L11 22.414V26h2v-7H6z"/>`),
		g.Group(children),
	)
}
package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Save(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.828a2 2 0 0 0-.586-1.414l-1.828-1.828A2 2 0 0 0 16.172 4H15M8 4v4a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1V4M8 4h7M7 17v-3a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v3"/>`),
		g.Group(children),
	)
}
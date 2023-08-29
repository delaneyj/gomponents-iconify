package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileZipLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 22H4a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1Zm-1-2V4H5v16h14Zm-5-8v5h-4v-3h2v-2h2Zm-2-8h2v2h-2V4Zm-2 2h2v2h-2V6Zm2 2h2v2h-2V8Zm-2 2h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}
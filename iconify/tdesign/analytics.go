package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analytics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h16V4H4Zm14 4v3h-2V8h2Zm-5 2v3h-2v-3h2Zm-5 2v6H6v-6h2Zm10 1v5h-2v-5h2Zm-5 2v3h-2v-3h2Z"/>`),
		g.Group(children),
	)
}
package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentDividerHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 2v20h-2V2h2ZM2 5h7v14H2V5Zm2 2v10h3V7H4Zm11-2h7v14h-7V5Zm2 2v10h3V7h-3Z"/>`),
		g.Group(children),
	)
}
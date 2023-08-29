package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentCheckbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h13v5.5h-2V4H4v9h3.5v2H2V2Zm7 7h13v13H9V9Zm2 2v9h9v-9h-9Zm8.414 3L15 18.414L12.086 15.5l1.414-1.414l1.5 1.5l3-3L19.414 14Z"/>`),
		g.Group(children),
	)
}
package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<circle cx="10" cy="10" r="1.75" fill="currentColor"/><path fill="currentColor" d="M15 1H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2zm0 9.75l-1.37.25a3.73 3.73 0 0 1-.38.93l.82 1.07L13 14.07l-1.12-.82a3.73 3.73 0 0 1-.93.38l-.2 1.37h-1.5L9 13.63a3.73 3.73 0 0 1-.93-.38L7 14.07L5.93 13l.82-1.12a3.73 3.73 0 0 1-.38-.88L5 10.75v-1.5L6.37 9a3.72 3.72 0 0 1 .38-.93L5.93 7L7 5.93l1.12.82A3.73 3.73 0 0 1 9 6.37L9.25 5h1.5L11 6.37a3.74 3.74 0 0 1 .93.38L13 5.93L14.07 7l-.82 1.12a3.73 3.73 0 0 1 .38.93l1.37.2z"/>`),
		g.Group(children),
	)
}
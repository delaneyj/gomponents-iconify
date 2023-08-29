package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicRectangleAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v11h-2V4H4v16h9v2H2V2Zm10 5h4v2h-2v5a3 3 0 1 1-2-2.83V7Zm0 7a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm8 1v3h3v2h-3v3h-2v-3h-3v-2h3v-3h2Z"/>`),
		g.Group(children),
	)
}
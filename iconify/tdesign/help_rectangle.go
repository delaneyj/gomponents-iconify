package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelpRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h16V4H4Zm8 3.25a2 2 0 0 0-1.886 1.333l-.334.943l-1.885-.666l.333-.943a4.001 4.001 0 1 1 6.246 4.476c-.431.34-.817.666-1.096 1.009c-.274.338-.378.61-.378.848v1.25h-2v-1.25c0-.867.39-1.573.826-2.11c.432-.53.974-.974 1.41-1.318A2 2 0 0 0 12 7.25Zm-1 9.25h2.004v2.004H11V16.5Z"/>`),
		g.Group(children),
	)
}
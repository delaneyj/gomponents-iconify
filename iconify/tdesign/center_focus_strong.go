package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CenterFocusStrong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h7v2H4v5H2V2Zm13 0h7v7h-2V4h-5V2Zm-3 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0Zm-4 3v5h5v2H2v-7h2Zm18 0v7h-7v-2h5v-5h2Z"/>`),
		g.Group(children),
	)
}
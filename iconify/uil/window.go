package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Window(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 5a1 1 0 1 0 1 1a1 1 0 0 0-1-1ZM6 5a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm8 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm6-4H4a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3Zm1 19a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-9h18Zm0-11H3V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}
package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 20h-1V8h1a1 1 0 0 0 0-2h-1V3a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v3H2a1 1 0 0 0 0 2h1v12H2a1 1 0 0 0 0 2h20a1 1 0 0 0 0-2Zm-7 0H9v-3.53a6.21 6.21 0 0 1 3-5.33a6.21 6.21 0 0 1 3 5.33Zm4 0h-2v-3.53a8.17 8.17 0 0 0-4.55-7.36a1 1 0 0 0-.9 0A8.17 8.17 0 0 0 7 16.47V20H5V8h14Zm0-14H5V4h14Z"/>`),
		g.Group(children),
	)
}
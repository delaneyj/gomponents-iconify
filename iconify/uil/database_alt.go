package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm8-9H8a4 4 0 0 0-4 4v12a4 4 0 0 0 4 4h8a4 4 0 0 0 4-4V6a4 4 0 0 0-4-4Zm2 16a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2v-2.56A3.91 3.91 0 0 0 8 16h8a3.91 3.91 0 0 0 2-.56Zm0-6a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V9.44A3.91 3.91 0 0 0 8 10h8a3.91 3.91 0 0 0 2-.56Zm-2-4H8a2 2 0 0 1 0-4h8a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}
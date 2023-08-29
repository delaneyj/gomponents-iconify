package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicSquareF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 0h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4zm6 3.605a3 3 0 0 0-3 3v4a3 3 0 1 0 6 0v-4a3 3 0 0 0-3-3zm-6 7v.022a6 6 0 1 0 12 0v-.022a1 1 0 0 0-2 0v.022a4 4 0 1 1-8 0v-.022a1 1 0 0 0-2 0zm6-5a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0v-4a1 1 0 0 1 1-1z"/>`),
		g.Group(children),
	)
}
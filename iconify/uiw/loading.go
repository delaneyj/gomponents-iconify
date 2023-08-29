package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loading(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 16a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm-6.259-3a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Zm11.578.5a2 2 0 1 1 0 4a2 2 0 0 1 0-4ZM18.5 9.319a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3ZM2.5 6a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Zm15.286-.793a1 1 0 1 1 0 2a1 1 0 0 1 0-2ZM8 0a3 3 0 1 1 0 6a3 3 0 0 1 0-6Zm7.5 3a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}
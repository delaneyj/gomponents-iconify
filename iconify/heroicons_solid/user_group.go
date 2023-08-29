package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 6a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm5 2a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm-4 7a4 4 0 0 0-8 0v3h8v-3ZM6 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm10 10v-3a5.972 5.972 0 0 0-.75-2.906A3.005 3.005 0 0 1 19 15v3h-3ZM4.75 12.094A5.973 5.973 0 0 0 4 15v3H1v-3a3 3 0 0 1 3.75-2.906Z"/>`),
		g.Group(children),
	)
}
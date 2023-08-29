package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.25 8A.75.75 0 0 1 9 7.25h7a.75.75 0 0 1 0 1.5H9A.75.75 0 0 1 8.25 8ZM9 10.25a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 0-1.5H9Z"/><path fill="currentColor" fill-rule="evenodd" d="M8.5 3.25A4.75 4.75 0 0 0 3.75 8v10a3.75 3.75 0 0 0 3.75 3.75h11A1.75 1.75 0 0 0 20.25 20V5a1.75 1.75 0 0 0-1.75-1.75h-10Zm10.25 11V5a.25.25 0 0 0-.25-.25h-10A3.25 3.25 0 0 0 5.25 8v7a3.734 3.734 0 0 1 2.25-.75h11.25Zm0 1.5H7.5a2.25 2.25 0 0 0 0 4.5h11a.25.25 0 0 0 .25-.25v-4.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
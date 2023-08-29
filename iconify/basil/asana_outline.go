package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsanaOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.47 3.554a4 4 0 1 1 3.06 7.391a4 4 0 0 1-3.06-7.39ZM12 4.75a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Zm-6.53 7.805a4 4 0 1 1 3.061 7.391a4 4 0 0 1-3.062-7.392ZM7 13.75a2.5 2.5 0 1 0 0 5.001a2.5 2.5 0 0 0 0-5.001Zm10-1.5a4 4 0 1 0 0 8a4 4 0 0 0 0-8Zm-.957 1.69a2.5 2.5 0 1 1 1.914 4.62a2.5 2.5 0 0 1-1.914-4.62Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
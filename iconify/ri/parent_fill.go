package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParentFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 11a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm10.5 4a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm0 1a4.5 4.5 0 0 1 4.5 4.5v.5h-9v-.5a4.5 4.5 0 0 1 4.5-4.5ZM7 12a5 5 0 0 1 5 5v4H2v-4a5 5 0 0 1 5-5Z"/>`),
		g.Group(children),
	)
}
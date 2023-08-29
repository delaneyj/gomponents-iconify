package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m7.536 8.657l2.828-2.83a1 1 0 0 1 1.414 1.416l-3.535 3.535a1 1 0 0 1-1.415.001l-2.12-2.12a1 1 0 1 1 1.413-1.415zM8 16A8 8 0 1 1 8 0a8 8 0 0 1 0 16zm0-2A6 6 0 1 0 8 2a6 6 0 0 0 0 12z"/>`),
		g.Group(children),
	)
}
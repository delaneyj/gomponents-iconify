package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EllipsisV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 12a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm0-6a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm2-4a2 2 0 1 0-4 0a2 2 0 0 0 4 0Z"/>`),
		g.Group(children),
	)
}
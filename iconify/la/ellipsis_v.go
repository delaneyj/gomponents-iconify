package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EllipsisV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 6a1.999 1.999 0 1 0 0 4a1.999 1.999 0 1 0 0-4zm0 8a1.999 1.999 0 1 0 0 4a1.999 1.999 0 1 0 0-4zm0 8a1.999 1.999 0 1 0 0 4a1.999 1.999 0 1 0 0-4z"/>`),
		g.Group(children),
	)
}
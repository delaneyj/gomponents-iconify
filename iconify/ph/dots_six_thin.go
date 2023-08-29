package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsSixThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M68 92a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm60-8a8 8 0 1 0 8 8a8 8 0 0 0-8-8Zm68 16a8 8 0 1 0-8-8a8 8 0 0 0 8 8ZM60 156a8 8 0 1 0 8 8a8 8 0 0 0-8-8Zm68 0a8 8 0 1 0 8 8a8 8 0 0 0-8-8Zm68 0a8 8 0 1 0 8 8a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListDashes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M88 64a8 8 0 0 1 8-8h120a8 8 0 0 1 0 16H96a8 8 0 0 1-8-8Zm128 56H96a8 8 0 0 0 0 16h120a8 8 0 0 0 0-16Zm0 64H96a8 8 0 0 0 0 16h120a8 8 0 0 0 0-16ZM56 56H40a8 8 0 0 0 0 16h16a8 8 0 0 0 0-16Zm0 64H40a8 8 0 0 0 0 16h16a8 8 0 0 0 0-16Zm0 64H40a8 8 0 0 0 0 16h16a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}
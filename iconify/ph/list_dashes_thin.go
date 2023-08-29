package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListDashesThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M92 64a4 4 0 0 1 4-4h120a4 4 0 0 1 0 8H96a4 4 0 0 1-4-4Zm124 60H96a4 4 0 0 0 0 8h120a4 4 0 0 0 0-8Zm0 64H96a4 4 0 0 0 0 8h120a4 4 0 0 0 0-8ZM56 60H40a4 4 0 0 0 0 8h16a4 4 0 0 0 0-8Zm0 64H40a4 4 0 0 0 0 8h16a4 4 0 0 0 0-8Zm0 64H40a4 4 0 0 0 0 8h16a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}
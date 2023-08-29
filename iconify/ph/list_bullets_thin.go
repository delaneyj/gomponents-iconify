package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListBulletsThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M84 64a4 4 0 0 1 4-4h128a4 4 0 0 1 0 8H88a4 4 0 0 1-4-4Zm132 60H88a4 4 0 0 0 0 8h128a4 4 0 0 0 0-8Zm0 64H88a4 4 0 0 0 0 8h128a4 4 0 0 0 0-8ZM44 120a8 8 0 1 0 8 8a8 8 0 0 0-8-8Zm0-64a8 8 0 1 0 8 8a8 8 0 0 0-8-8Zm0 128a8 8 0 1 0 8 8a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListBulletsLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M82 64a6 6 0 0 1 6-6h128a6 6 0 0 1 0 12H88a6 6 0 0 1-6-6Zm134 58H88a6 6 0 0 0 0 12h128a6 6 0 0 0 0-12Zm0 64H88a6 6 0 0 0 0 12h128a6 6 0 0 0 0-12ZM44 54a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm0 128a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm0-64a10 10 0 1 0 10 10a10 10 0 0 0-10-10Z"/>`),
		g.Group(children),
	)
}
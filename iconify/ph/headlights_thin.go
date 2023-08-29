package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadlightsThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M164 80a4 4 0 0 1 4-4h72a4 4 0 0 1 0 8h-72a4 4 0 0 1-4-4Zm76 92h-72a4 4 0 0 0 0 8h72a4 4 0 0 0 0-8Zm0-64h-72a4 4 0 0 0 0 8h72a4 4 0 0 0 0-8Zm0 32h-72a4 4 0 0 0 0 8h72a4 4 0 0 0 0-8ZM140 64v128a12 12 0 0 1-12 12H88a76 76 0 0 1-76-76.58C12.31 85.83 46.81 52 88.9 52H128a12 12 0 0 1 12 12Zm-8 0a4 4 0 0 0-4-4H88.9C51.19 60 20.28 90.27 20 127.48A68 68 0 0 0 88 196h40a4 4 0 0 0 4-4Z"/>`),
		g.Group(children),
	)
}
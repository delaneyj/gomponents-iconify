package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SafeTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.005 20h-4v2h-2v-2h-1a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h7V1.59a.5.5 0 0 1 .582-.493L21.17 2.86a1 1 0 0 1 .836.987V6h1v2h-1v7h1v2h-1v2.152a1 1 0 0 1-.836.987l-1.164.194V22h-2v-1.334l-7.418 1.237a.5.5 0 0 1-.582-.494V20Zm2-.361l8-1.334V4.694l-8-1.333v16.278Zm4.5-5.64c-.828 0-1.5-1.119-1.5-2.5c0-1.38.671-2.5 1.5-2.5c.828 0 1.5 1.12 1.5 2.5c0 1.381-.672 2.5-1.5 2.5Z"/>`),
		g.Group(children),
	)
}
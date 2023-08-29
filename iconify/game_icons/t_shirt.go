package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TShirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m32 144l48 64l64-32l-16 304c64 16 192 16 256 0l-16-304l64 32l48-64l-112-96l-48-16c-16 64-112 64-128 0l-48 16z"/>`),
		g.Group(children),
	)
}
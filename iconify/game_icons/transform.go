package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M128 32a96 96 0 0 0-96 96a96 96 0 0 0 96 96a96 96 0 0 0 30.285-4.986L140.29 201.02l64.353-64.352l6.363-6.363l11.86 11.86A96 96 0 0 0 224 128a96 96 0 0 0-96-96zm83.006 123.76l-45.26 45.26L252.73 288l-23.468 23.467l115.24 23.047l-23.05-115.24l-23.466 23.466l-86.98-86.98zM353.556 288l13.89 69.46L288 341.57V480h192V288H353.555z"/>`),
		g.Group(children),
	)
}
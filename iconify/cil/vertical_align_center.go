package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalAlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M240 16v146.177l-53.491-53.49l-22.627 22.626L256 223.431l92.118-92.118l-22.627-22.626L272 162.177V16h-32zM16 240h480v32H16zm147.882 140.687l22.627 22.626L240 349.823V496h32V349.823l53.491 53.49l22.627-22.626L256 288.569l-92.118 92.118z"/>`),
		g.Group(children),
	)
}
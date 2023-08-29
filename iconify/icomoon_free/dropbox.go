package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11.5.5L8 3.5l4.5 3l3.5-3zM8 3.5L4.5.5L0 3.5l3.5 3zm4.5 3l3.5 3l-4.5 2.5L8 9zM8 9L3.5 6.5L0 9.5L4.5 12z"/><path fill="currentColor" d="M11.377 13.212L8 10.317l-3.377 2.895L2.5 12.033V13.5L8 16l5.5-2.5v-1.467z"/>`),
		g.Group(children),
	)
}
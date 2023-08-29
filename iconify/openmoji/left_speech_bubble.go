package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftSpeechBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FFF" d="M54.88 49.128A22.887 22.887 0 0 0 59 36c0-12.703-10.297-23-23-23S13 23.297 13 36s10.297 23 23 23c3.758 0 7.302-.907 10.435-2.505l4.814 2.052l5.728 2.443l-1.084-6.132l-1.012-5.73z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M54.88 49.128A22.887 22.887 0 0 0 59 36c0-12.703-10.297-23-23-23S13 23.297 13 36s10.297 23 23 23c3.758 0 7.302-.907 10.435-2.505l4.814 2.052l5.728 2.443l-1.084-6.132l-1.012-5.73z"/>`),
		g.Group(children),
	)
}
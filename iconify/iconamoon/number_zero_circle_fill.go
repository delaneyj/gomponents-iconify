package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberZeroCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm8-1.875C10 8.913 10.934 8 12 8c1.066 0 2 .913 2 2.125v3.75C14 15.088 13.066 16 12 16c-1.066 0-2-.912-2-2.125v-3.75ZM12 6c-2.247 0-4 1.886-4 4.125v3.75C8 16.115 9.753 18 12 18s4-1.886 4-4.125v-3.75C16 7.885 14.247 6 12 6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
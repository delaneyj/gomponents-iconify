package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquarePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.664 0H1.336C.598 0 .002.598.002 1.335v13.329C.002 15.401.598 16 1.336 16h13.328c.738 0 1.336-.599 1.336-1.336V1.335C16 .598 15.402 0 14.664 0zM9.023 9.016v2.55c0 .631-2.025.631-2.025 0v-2.55H4.421c-.633 0-.632-2.01 0-2.01h2.577V4.461c0-.633 2.025-.633 2.025 0v2.545h2.556c.633 0 .633 2.01 0 2.01H9.023z"/>`),
		g.Group(children),
	)
}
package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M24 511q-40-37-13-62L474 14q15-14 37.5-14T550 14l463 435q27 26-13 62H24zm104 256h768q53 0 90.5 37.5T1024 895t-37.5 90.5T896 1023H128q-53 0-90.5-37.5T0 895t37.5-90.5T128 767z"/>`),
		g.Group(children),
	)
}
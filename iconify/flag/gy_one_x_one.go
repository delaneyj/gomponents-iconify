package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GyOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#399408" d="M2 0h510v512H2z"/><path fill="#fff" d="M.1 0c-.6 0 495.7 257.6 495.7 257.6L0 511.7L.1 0z"/><path fill="#ffde08" d="M.2 21.5C3 21.5 447.5 254 445 256.2L1.5 494.2L.2 21.4z"/><path d="M1.5.8c1.5 0 232.8 257 232.8 257L1.5 508.8V.8z"/><path fill="#de2110" d="M.2 36.2C1.6 20.2 209 258.5 209 258.5L.2 481.8V36.2z"/></g>`),
		g.Group(children),
	)
}
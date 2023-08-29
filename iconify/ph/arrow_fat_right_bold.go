package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatRightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m240.49 119.51l-96-96A12 12 0 0 0 124 32v36H48a20 20 0 0 0-20 20v80a20 20 0 0 0 20 20h76v36a12 12 0 0 0 20.49 8.49l96-96a12 12 0 0 0 0-16.98ZM148 195v-19a12 12 0 0 0-12-12H52V92h84a12 12 0 0 0 12-12V61l67 67Z"/>`),
		g.Group(children),
	)
}
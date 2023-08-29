package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m232.48 111.51l-96-96a12 12 0 0 0-17 0l-96 96A12 12 0 0 0 32 132h36v76a20 20 0 0 0 20 20h80a20 20 0 0 0 20-20v-76h36a12 12 0 0 0 8.48-20.49ZM176 108a12 12 0 0 0-12 12v84H92v-84a12 12 0 0 0-12-12H61l67-67l67 67Z"/>`),
		g.Group(children),
	)
}
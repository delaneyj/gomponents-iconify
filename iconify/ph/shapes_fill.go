package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapesFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M111.59 181.47A8 8 0 0 1 104 192H24a8 8 0 0 1-7.59-10.53l40-120a8 8 0 0 1 15.18 0ZM208 76a52 52 0 1 0-52 52a52.06 52.06 0 0 0 52-52Zm16 68h-88a8 8 0 0 0-8 8v56a8 8 0 0 0 8 8h88a8 8 0 0 0 8-8v-56a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}
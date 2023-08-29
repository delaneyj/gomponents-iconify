package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleDownDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m208 48l-80 80l-80-80Z" opacity=".2"/><path d="m213.66 133.66l-80 80a8 8 0 0 1-11.32 0l-80-80a8 8 0 0 1 11.32-11.32L128 196.69l74.34-74.35a8 8 0 0 1 11.32 11.32Zm-171.32-80A8 8 0 0 1 48 40h160a8 8 0 0 1 5.66 13.66l-80 80a8 8 0 0 1-11.32 0Zm25 2.34L128 116.69L188.69 56Z"/></g>`),
		g.Group(children),
	)
}
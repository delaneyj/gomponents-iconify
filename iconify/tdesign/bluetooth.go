package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 .255l7.453 6.707L13.414 12l5.039 5.038L11 23.745v-9.33l-4 4L5.586 17l5-5l-5-5L7 5.586l4 4V.255Zm2 14.16v4.84l2.548-2.293L13 14.414Zm0-4.83l2.548-2.547L13 4.745v4.84Z"/>`),
		g.Group(children),
	)
}
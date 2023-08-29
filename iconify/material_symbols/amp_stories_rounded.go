package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmpStoriesRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 20q-.425 0-.713-.288T7 19V5q0-.425.288-.713T8 4h8q.425 0 .713.288T17 5v14q0 .425-.288.713T16 20H8Zm-5-3V6.975q0-.425.288-.7T4 6q.425 0 .713.288T5 7v10.025q0 .425-.288.7T4 18q-.425 0-.713-.288T3 17Zm16 0V6.975q0-.425.288-.7T20 6q.425 0 .713.288T21 7v10.025q0 .425-.288.7T20 18q-.425 0-.713-.288T19 17Z"/>`),
		g.Group(children),
	)
}
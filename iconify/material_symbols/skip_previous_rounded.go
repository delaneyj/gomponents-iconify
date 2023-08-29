package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipPreviousRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 18q-.425 0-.713-.288T5.5 17V7q0-.425.288-.713T6.5 6q.425 0 .713.288T7.5 7v10q0 .425-.288.713T6.5 18Zm10.45-1.025l-6.2-4.15q-.45-.3-.45-.825t.45-.825l6.2-4.15q.5-.325 1.025-.037t.525.887v8.25q0 .6-.525.9t-1.025-.05Z"/>`),
		g.Group(children),
	)
}
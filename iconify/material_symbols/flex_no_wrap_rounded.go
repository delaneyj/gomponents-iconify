package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlexNoWrapRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 16V8q0-.425.288-.713T2 7h4q.425 0 .713.288T7 8v8q0 .425-.288.713T6 17H2q-.425 0-.713-.288T1 16Zm8 0V8q0-.425.288-.713T10 7h4q.425 0 .713.288T15 8v8q0 .425-.288.713T14 17h-4q-.425 0-.713-.288T9 16Zm8 0V8q0-.425.288-.713T18 7h4q.425 0 .713.288T23 8v8q0 .425-.288.713T22 17h-4q-.425 0-.713-.288T17 16ZM3 15h2V9H3v6Zm16 0h2V9h-2v6Z"/>`),
		g.Group(children),
	)
}
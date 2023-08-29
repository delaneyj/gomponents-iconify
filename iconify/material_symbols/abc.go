package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Abc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 15q-.425 0-.713-.288T16 14v-4q0-.425.288-.713T17 9h3q.425 0 .713.288T21 10v1h-1.5v-.5h-2v3h2V13H21v1q0 .425-.288.713T20 15h-3Zm-7.5 0V9h4q.425 0 .713.288T14.5 10v1q0 .425-.288.713T13.5 12q.425 0 .713.288T14.5 13v1q0 .425-.288.713T13.5 15h-4Zm1.5-3.75h2v-.75h-2v.75Zm0 2.25h2v-.75h-2v.75ZM3 15v-5q0-.425.288-.713T4 9h3q.425 0 .713.288T8 10v5H6.5v-1.5h-2V15H3Zm1.5-3h2v-1.5h-2V12Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalSplitOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15q-.425 0-.713-.288T3 14q0-.425.275-.713t.7-.287H10q.425 0 .713.288T11 14q0 .425-.275.713t-.7.287H4Zm0 4q-.425 0-.713-.288T3 18q0-.425.275-.713t.7-.287H10q.425 0 .713.288T11 18q0 .425-.275.713t-.7.287H4Zm0-8q-.425 0-.713-.288T3 10q0-.425.275-.713t.7-.287H10q.425 0 .713.288T11 10q0 .425-.275.713t-.7.287H4Zm0-4q-.425 0-.713-.288T3 6q0-.425.275-.713t.7-.287H10q.425 0 .713.288T11 6q0 .425-.275.713t-.7.287H4Zm11 0v10V7Zm-1 12q-.425 0-.713-.288T13 18V6q0-.425.288-.713T14 5h6q.425 0 .713.288T21 6v12q0 .425-.288.713T20 19h-6Zm1-12v10h4V7h-4Z"/>`),
		g.Group(children),
	)
}
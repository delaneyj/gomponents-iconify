package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewHeadlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.975 15q-.425 0-.7-.288T4 14q0-.425.288-.713T5 13h14.025q.425 0 .7.288T20 14q0 .425-.288.713T19 15H4.975Zm0 4q-.425 0-.7-.288T4 18q0-.425.288-.713T5 17h14.025q.425 0 .7.288T20 18q0 .425-.288.713T19 19H4.975Zm0-8q-.425 0-.7-.288T4 10q0-.425.288-.713T5 9h14.025q.425 0 .7.288T20 10q0 .425-.288.713T19 11H4.975Zm0-4q-.425 0-.7-.288T4 6q0-.425.288-.713T5 5h14.025q.425 0 .7.288T20 6q0 .425-.288.713T19 7H4.975Z"/>`),
		g.Group(children),
	)
}
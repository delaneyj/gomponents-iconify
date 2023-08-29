package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FenceRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 20q-.425 0-.713-.288T5 19v-3H4q-.425 0-.713-.288T3 15q0-.425.288-.713T4 14h1v-2H4q-.425 0-.713-.288T3 11q0-.425.288-.713T4 10h1V7.425q0-.2.075-.387T5.3 6.7l2-2q.3-.3.7-.3t.7.3L10 6l1.3-1.3q.3-.3.713-.3t.712.3l1.3 1.3L15.3 4.7q.3-.3.713-.3t.712.3l2 2q.15.15.225.338t.075.387V10H20q.425 0 .713.288T21 11q0 .425-.288.713T20 12h-.975v2H20q.425 0 .713.288T21 15q0 .425-.288.713T20 16h-.975v3q0 .425-.288.713t-.712.287H6Zm1-10h2V7.825l-1-1l-1 1V10Zm4 0h2V7.825l-1-1l-1 1V10Zm4.025 0H17V7.825l-1-1l-.975.975V10ZM7 14h2v-2H7v2Zm4 0h2v-2h-2v2Zm4.025 0H17v-2h-1.975v2ZM7 18h2v-2H7v2Zm4 0h2v-2h-2v2Zm4.025 0H17v-2h-1.975v2Z"/>`),
		g.Group(children),
	)
}
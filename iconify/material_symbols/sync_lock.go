package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyncLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20v-2h2.725q-1.275-1.1-2-2.65T4 12q0-2.8 1.7-4.938T10 4.25v2.1q-1.75.625-2.875 2.163T6 12q0 1.35.537 2.488T8 16.45V14h2v6H4Zm14-8q0-1.275-.513-2.388T16 7.55V10h-2V4h6v2h-2.725q1.475 1.325 2.088 2.838T20 12h-2Zm-2 10q-.425 0-.713-.288T15 21v-3q0-.425.288-.713T16 17v-1q0-.825.588-1.413T18 14q.825 0 1.413.588T20 16v1q.425 0 .713.288T21 18v3q0 .425-.288.713T20 22h-4Zm1-5h2v-1q0-.425-.288-.713T18 15q-.425 0-.713.288T17 16v1Z"/>`),
		g.Group(children),
	)
}
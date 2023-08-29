package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MosqueRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.95 8.05q-.525 0-.738-.213T6 7.1q0-1 .475-1.85t1.275-1.4l3.125-2.1q.5-.35 1.125-.35t1.125.35l3.125 2.1q.8.55 1.275 1.4T18 7.1q0 .525-.213.738t-.737.212H6.95ZM1 19V8.725Q.55 8.45.275 8.012T0 7q0-.575.6-1.4T2 4q.8.775 1.4 1.6T4 7q0 .575-.275 1.012T3 8.725V13h2v-2q0-.625.4-1.2t1.15-.75h10.9q.75.175 1.15.75T19 11v2h2V8.725q-.45-.275-.725-.713T20 7q0-.575.6-1.4T22 4q.8.775 1.4 1.6T24 7q0 .575-.275 1.012T23 8.725V19q0 .825-.588 1.413T21 21h-6q-.425 0-.713-.288T14 20v-3q0-.825-.588-1.413T12 15q-.825 0-1.413.588T10 17v3q0 .425-.288.713T9 21H3q-.825 0-1.413-.588T1 19Z"/>`),
		g.Group(children),
	)
}
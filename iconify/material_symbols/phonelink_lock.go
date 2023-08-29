package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonelinkLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.85 16q-.35 0-.6-.25t-.25-.6v-3.3q0-.35.25-.6t.6-.25H16v-1q0-.825.588-1.412T18 8q.825 0 1.413.588T20 10v1h.15q.35 0 .6.25t.25.6v3.3q0 .35-.25.6t-.6.25h-4.3ZM17 11h2v-1q0-.425-.288-.713T18 9q-.425 0-.713.288T17 10v1ZM7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v4h-2V6H7v12h10v-1h2v4q0 .825-.588 1.413T17 23H7Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintLockSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 21v-5h1v-1q0-.825.588-1.413T19 13q.825 0 1.413.588T21 15v1h1v5h-6Zm2-5h2v-1q0-.425-.288-.713T19 14q-.425 0-.713.288T18 15v1ZM6 21v-4H2V8h20v3.75q-.675-.35-1.413-.55t-1.512-.2q-1.95 0-3.538 1.1T13.25 15H8v4h5.1q.175.55.425 1.05t.6.95H6ZM6 7V3h12v4H6Z"/>`),
		g.Group(children),
	)
}
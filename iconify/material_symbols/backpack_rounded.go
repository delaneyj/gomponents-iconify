package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackpackRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20V8q0-1.4.85-2.45T7 4.15V3q0-.425.288-.713T8 2h1q.425 0 .713.288T10 3v1h4V3q0-.425.288-.713T15 2h1q.425 0 .713.288T17 3v1.15q1.3.35 2.15 1.4T20 8v12q0 .825-.588 1.413T18 22H6Zm8.5-8v1q0 .425.288.713T15.5 16q.425 0 .713-.288T16.5 15v-2q0-.425-.288-.713T15.5 12h-7q-.425 0-.713.288T7.5 13q0 .425.288.713T8.5 14h6Z"/>`),
		g.Group(children),
	)
}
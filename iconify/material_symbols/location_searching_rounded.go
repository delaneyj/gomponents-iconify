package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationSearchingRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22v-1q-3.125-.35-5.363-2.588T3.05 13.05h-1q-.425 0-.712-.288t-.288-.712q0-.425.288-.713t.712-.287h1q.35-3.125 2.588-5.363T11 3.1v-1q0-.425.288-.713T12 1.1q.425 0 .713.288T13 2.1v1q3.125.35 5.363 2.588t2.587 5.362h1q.425 0 .713.288t.287.712q0 .425-.287.713t-.713.287h-1q-.35 3.125-2.587 5.363T13 21v1q0 .425-.288.713T12 23q-.425 0-.713-.288T11 22Zm1-2.95q2.9 0 4.95-2.05T19 12.05q0-2.9-2.05-4.95T12 5.05q-2.9 0-4.95 2.05T5 12.05q0 2.9 2.05 4.95T12 19.05Z"/>`),
		g.Group(children),
	)
}
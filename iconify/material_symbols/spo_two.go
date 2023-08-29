package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpoTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q-.425 0-.713-.288T11 19v-4q0-.425.288-.713T12 14h2.5q.425 0 .713.288T15.5 15v4q0 .425-.288.713T14.5 20H12Zm.5-1.5H14v-3h-1.5v3ZM17 22v-2.75q0-.425.288-.713T18 18.25h2v-.75h-3V16h3.5q.425 0 .713.288T21.5 17v1.75q0 .425-.288.713t-.712.287h-2v.75h3V22H17Zm-8-.05q-3.075-.35-5.038-2.638T2 13.8q0-2.5 1.988-5.438T10 2q3.3 2.8 5.238 5.3T17.75 12H11q-.825 0-1.413.588T9 14v7.95Z"/>`),
		g.Group(children),
	)
}
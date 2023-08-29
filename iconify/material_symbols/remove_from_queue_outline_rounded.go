package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveFromQueueOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 12h6q.425 0 .713-.288T16 11q0-.425-.288-.713T15 10H9q-.425 0-.713.288T8 11q0 .425.288.713T9 12Zm-5 7q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v12q0 .825-.588 1.413T20 19h-4v1q0 .425-.288.713T15 21H9q-.425 0-.713-.288T8 20v-1H4Zm0-2h16V5H4v12Zm0 0V5v12Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtitlesOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 10h-5.15l-6-6H20q.825 0 1.413.588T22 6v11.9q0 .275-.05.525t-.2.475l-6.9-6.9H18v-2Zm2.55 13.35L17.15 20H4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4v2.8L.65 3.45l1.4-1.4l19.9 19.9l-1.4 1.4ZM6 12h2v-1.2l-.8-.8H6v2Zm5.15 2H6v2h7.15l-2-2Z"/>`),
		g.Group(children),
	)
}
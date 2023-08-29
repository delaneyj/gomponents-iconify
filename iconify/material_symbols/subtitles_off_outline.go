package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtitlesOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.55 23.35L17.15 20H4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4l2 2H4v12h11.15l-2-2H6v-2h5.15L.65 3.45l1.4-1.4l19.9 19.9l-1.4 1.4Zm1.2-4.45L20 17.15V6H8.85l-2-2H20q.825 0 1.413.588T22 6v11.9q0 .275-.05.525t-.2.475Zm-6.9-6.9l-2-2H18v2h-3.15ZM6 12v-2h2v2H6Zm8.425-.425Zm-4.85.85Z"/>`),
		g.Group(children),
	)
}
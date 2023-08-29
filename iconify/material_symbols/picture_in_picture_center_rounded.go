package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureInPictureCenterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm5-5h6q.425 0 .713-.288T16 14v-4q0-.425-.288-.713T15 9H9q-.425 0-.713.288T8 10v4q0 .425.288.713T9 15Z"/>`),
		g.Group(children),
	)
}
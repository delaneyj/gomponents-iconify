package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartnerReportsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q-.425 0-.713-.288T11 16q0-.425.288-.713T12 15q.425 0 .713.288T13 16q0 .425-.288.713T12 17Zm0-4q-.425 0-.713-.288T11 12V4q0-.425.288-.713T12 3q.425 0 .713.288T13 4v8q0 .425-.288.713T12 13Zm-7 8q-.825 0-1.413-.588T3 19v-2q0-.425.288-.713T4 16q.425 0 .713.288T5 17v2h14v-2q0-.425.288-.713T20 16q.425 0 .713.288T21 17v2q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}
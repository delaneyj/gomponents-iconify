package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartnerReports(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q-.425 0-.713-.288T11 16q0-.425.288-.713T12 15q.425 0 .713.288T13 16q0 .425-.288.713T12 17Zm-1-4V3h2v10h-2Zm-6 8q-.825 0-1.413-.588T3 19v-3h2v3h14v-3h2v3q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}
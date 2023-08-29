package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDownloadOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.075 21.9L17.15 20H6q-.825 0-1.413-.588T4 18v-2q0-.425.288-.713T5 15q.425 0 .713.288T6 16v2h9.15l-2.575-2.575q-.15.075-.287.113t-.288.037q-.25 0-.413-.075t-.287-.2l-3.6-3.6q-.275-.275-.288-.637t.188-.613L2.075 4.925q-.275-.275-.275-.7t.3-.725q.275-.275.7-.275t.725.275l16.975 17q.275.275.275.7t-.275.7q-.3.3-.725.3t-.7-.3ZM20 17.15l-1.775-1.775q.125-.15.313-.262T19 15q.425 0 .713.288T20 16v1.15Zm-4.575-4.575L14 11.15l.875-.875q.275-.275.713-.263t.712.288q.275.275.275.7t-.275.7l-.875.875ZM13 10.15l-2-2V5q0-.425.288-.713T12 4q.425 0 .713.288T13 5v5.15Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PriceCheckRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 14.975q-.425 0-.713-.287t-.287-.713H5q-.425 0-.713-.287T4 12.975q0-.425.288-.712T5 11.975h4v-2H5q-.425 0-.713-.287T4 8.975v-4q0-.425.288-.712T5 3.975h1.5q0-.425.288-.713t.712-.287q.425 0 .713.288t.287.712H10q.425 0 .713.288t.287.712q0 .425-.288.713T10 5.975H6v2h4q.425 0 .713.288t.287.712v4q0 .425-.288.713t-.712.287H8.5q0 .425-.288.713t-.712.287Zm6.45 5.575q-.2 0-.375-.063t-.325-.212l-2.85-2.85q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l2.15 2.15l4.95-4.95q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-5.65 5.65q-.15.15-.325.213t-.375.062Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareReviewsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 15q.825 0 1.413-.588T16.5 13q0-.825-.588-1.413T14.5 11q-.375 0-.713.138t-.612.362L10.5 10.15v-.3l2.675-1.35q.275.225.613.363T14.5 9q.825 0 1.413-.588T16.5 7q0-.825-.588-1.413T14.5 5q-.825 0-1.413.588T12.5 7v.15L9.825 8.5q-.275-.225-.613-.363T8.5 8q-.825 0-1.413.588T6.5 10q0 .825.588 1.413T8.5 12q.375 0 .713-.138t.612-.362l2.675 1.35V13q0 .825.588 1.413T14.5 15ZM2 22V2h20v16H6l-4 4Zm2-6h16V4H4v12Zm0 0V4v12Z"/>`),
		g.Group(children),
	)
}
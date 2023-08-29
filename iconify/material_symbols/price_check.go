package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PriceCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 14.975v-1H4v-2h5v-2H5q-.425 0-.713-.287T4 8.975v-4q0-.425.288-.712T5 3.975h1.5v-1h2v1H11v2H6v2h4q.425 0 .713.288t.287.712v4q0 .425-.288.713t-.712.287H8.5v1h-2Zm7.45 6l-4.25-4.25l1.4-1.4l2.85 2.85l5.65-5.65l1.4 1.4l-7.05 7.05Z"/>`),
		g.Group(children),
	)
}
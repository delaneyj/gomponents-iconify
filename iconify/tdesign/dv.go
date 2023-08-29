package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h10v6.682c.456.687.775 1.473.917 2.318H14V7h8v10h-8v-4h-1.083A5.972 5.972 0 0 1 12 15.318V22H2v-6.682c-.632-.95-1-2.093-1-3.318s.368-2.367 1-3.318V2Zm2 4.803A5.973 5.973 0 0 1 7 6c1.093 0 2.118.293 3 .803V4H4v2.803Zm0 10.394V20h6v-2.803A5.974 5.974 0 0 1 7 18a5.974 5.974 0 0 1-3-.803ZM7 8a3.992 3.992 0 0 0-3.2 1.6A3.978 3.978 0 0 0 3 12c0 .902.297 1.731.8 2.4A3.992 3.992 0 0 0 7 16a3.992 3.992 0 0 0 3.2-1.6c.503-.669.8-1.498.8-2.4c0-.902-.297-1.731-.8-2.4A3.992 3.992 0 0 0 7 8Zm9 1v6h4V9h-4Z"/>`),
		g.Group(children),
	)
}
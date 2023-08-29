package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiChannel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q.35-2.9.85-5.288t1.137-4.1q.638-1.712 1.4-2.662T8 8q1 0 1.975 1.425t1.625 4.15q.325-2.55.763-4.525t.987-3.325q.55-1.35 1.213-2.037T16 3q1.075 0 1.925 1.275T19.4 7.9q.625 2.35 1.025 5.675T21 21h-2q-.45-2.2-1.4-4.1T16 15q-.65 0-1.613 1.9T13 21h-2q-.2-1.775-.537-3.538t-.75-3.237q-.413-1.475-.85-2.563T8 10.125q-.4.45-.85 1.525t-.863 2.55q-.412 1.475-.75 3.238T5 21H3Zm10.5-6.5q.575-.75 1.2-1.125T16 13q.65 0 1.288.388T18.5 14.5q-.45-3.8-1.125-6.338T16 5.05q-.7.575-1.375 3.15T13.5 14.5Z"/>`),
		g.Group(children),
	)
}
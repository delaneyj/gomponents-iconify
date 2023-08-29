package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BikePedal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.85 5c-.98 0-1.81.71-1.98 1.67L4 12l.87 5.33c.17.96 1 1.67 1.98 1.67h7.3c.98 0 1.82-.71 1.98-1.67l.55-3.33H18v-1h2v-2h-2v-1h-1.32l-.55-3.33c-.16-.96-1-1.67-1.98-1.67h-7.3m0 2h7.3l.67 4H6.18l.67-4m-.67 6h8.64l-.67 4h-7.3l-.67-4Z"/>`),
		g.Group(children),
	)
}
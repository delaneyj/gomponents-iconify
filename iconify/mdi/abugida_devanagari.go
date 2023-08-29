package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbugidaDevanagari(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 3v2h3c1.32 0 2.41.83 2.82 2H6v2h8v1h-2c-2.75 0-5 2.25-5 5s2.25 5 5 5c.77 0 1.45-.27 2-.7V21h2v-4h-2c-.45.62-1.17 1-2 1c-1.67 0-3-1.33-3-3s1.33-3 3-3h4V9h2V7h-2.1c-.47-2.28-2.49-4-4.9-4H8Z"/>`),
		g.Group(children),
	)
}
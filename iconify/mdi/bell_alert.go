package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 7v6h-2V7m0 8h2v2h-2M12 2a2 2 0 0 0-2 2a2 2 0 0 0 0 .29C7.12 5.14 5 7.82 5 11v6l-2 2v1h18v-1l-2-2v-6c0-3.18-2.12-5.86-5-6.71A2 2 0 0 0 14 4a2 2 0 0 0-2-2m-2 19a2 2 0 0 0 2 2a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}
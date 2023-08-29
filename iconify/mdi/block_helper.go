package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockHelper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0a12 12 0 0 1 12 12a12 12 0 0 1-12 12A12 12 0 0 1 0 12A12 12 0 0 1 12 0m0 2A10 10 0 0 0 2 12c0 2.4.85 4.6 2.26 6.33L18.33 4.26A9.984 9.984 0 0 0 12 2m0 20a10 10 0 0 0 10-10c0-2.4-.85-4.6-2.26-6.33L5.67 19.74A9.984 9.984 0 0 0 12 22Z"/>`),
		g.Group(children),
	)
}
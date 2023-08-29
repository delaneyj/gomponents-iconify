package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.5 12c0-3.9.8-6.9 1.4-8.7c.2-.6-.2-1.3-.9-1.3H5c-.7 0-1.2.7-.9 1.3c.6 2.1 1.4 4.8 1.4 8.7c0 3.9-.8 6.7-1.4 8.7c-.3.6.2 1.3.9 1.3h14c.7 0 1.2-.7 1-1.3c-.7-2-1.5-4.8-1.5-8.7Z"/>`),
		g.Group(children),
	)
}
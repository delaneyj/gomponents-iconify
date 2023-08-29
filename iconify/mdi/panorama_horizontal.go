package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 5.5c-3.9 0-6.9-.8-8.7-1.4c-.6-.3-1.3.2-1.3.9v14c0 .7.7 1.2 1.3 1c2.1-.7 4.8-1.5 8.7-1.5c3.9 0 6.7.8 8.7 1.5c.7.2 1.3-.3 1.3-1V5c0-.7-.7-1.2-1.3-.9c-2 .6-4.8 1.4-8.7 1.4Z"/>`),
		g.Group(children),
	)
}
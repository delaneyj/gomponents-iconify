package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOutTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 8a.5.5 0 0 1 0 1H6a.5.5 0 0 1 0-1h5Zm3 .5a5.5 5.5 0 1 0-1.98 4.227l4.126 4.127l.07.057a.5.5 0 0 0 .638-.765l-4.127-4.126A5.478 5.478 0 0 0 14 8.5Zm-10 0a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0Z"/>`),
		g.Group(children),
	)
}
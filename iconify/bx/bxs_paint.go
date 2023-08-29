package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsPaint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M21.084 2.914c-1.178-1.179-3.234-1.179-4.412 0l-8.379 8.379a.999.999 0 0 0 0 1.414l3 3a.997.997 0 0 0 1.414 0l8.379-8.379a3.123 3.123 0 0 0-.002-4.414zm-1.412 3L12 13.586L10.414 12l7.672-7.672a1.146 1.146 0 0 1 1.586.002a1.123 1.123 0 0 1 0 1.584zM8 15c-1.265-.634-3.5 0-3.5 2c0 1.197.5 2-1.5 3c0 0 3.25 2.25 5.5 0c1.274-1.274 1.494-4-.5-5z" fill="currentColor"/>`),
		g.Group(children),
	)
}
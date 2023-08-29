package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileRotateRightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 11h3l-4 4l-4-4h3c0-3.31-2.69-6-6-6l-1 .08V3.06L12 3c4.42 0 8 3.58 8 8M9 7H5c-1.1 0-2 .9-2 2v9a2 2 0 0 0 2 2h6c1.11 0 2-.89 2-2v-7L9 7m2 11H5V9h3v3h3v6Z"/>`),
		g.Group(children),
	)
}
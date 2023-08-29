package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteRestore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 14h2l-4-4l-4 4h2v4h4v-4M6 7h12v12c0 .5-.2 1-.61 1.39c-.39.41-.89.61-1.39.61H8c-.5 0-1-.2-1.39-.61C6.2 20 6 19.5 6 19V7m13-3v2H5V4h3.5l1-1h5l1 1H19Z"/>`),
		g.Group(children),
	)
}
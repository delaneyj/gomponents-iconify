package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.527 0H6.24l10.239 16L6.24 32h9.287L25.76 16L15.527 0z"/>`),
		g.Group(children),
	)
}
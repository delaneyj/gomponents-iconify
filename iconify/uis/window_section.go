package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowSection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21c0 .6.4 1 1 1h5V10H2v11zm14 1h5c.6 0 1-.4 1-1V10h-6v12zm-6 0h4V10h-4v12zM21 2H3c-.6 0-1 .4-1 1v5h20V3c0-.6-.4-1-1-1z"/>`),
		g.Group(children),
	)
}
package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoltAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.89 9.55A1 1 0 0 0 19 9h-5V3a1 1 0 0 0-.69-1a1 1 0 0 0-1.12.36l-8 11a1 1 0 0 0-.08 1A1 1 0 0 0 5 15h5v6a1 1 0 0 0 .69.95A1.12 1.12 0 0 0 11 22a1 1 0 0 0 .81-.41l8-11a1 1 0 0 0 .08-1.04ZM12 17.92V14a1 1 0 0 0-1-1H7l5-6.92V10a1 1 0 0 0 1 1h4Z"/>`),
		g.Group(children),
	)
}
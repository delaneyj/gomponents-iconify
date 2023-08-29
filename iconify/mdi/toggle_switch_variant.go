package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleSwitchVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.4 1.6C18 1.2 17.5 1 17 1H7c-.5 0-1 .2-1.4.6C5.2 2 5 2.5 5 3v18c0 .5.2 1 .6 1.4c.4.4.9.6 1.4.6h10c.5 0 1-.2 1.4-.6c.4-.4.6-.9.6-1.4V3c0-.5-.2-1-.6-1.4M16 7c0 .6-.4 1-1 1H9c-.6 0-1-.4-1-1V5c0-.6.4-1 1-1h6c.6 0 1 .4 1 1v2Z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightSwitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 6v12h8V6H8m6 4h-4V8h4v2m5.4-8.4C19 1.2 18.5 1 18 1H6c-.5 0-1 .2-1.4.6C4.2 2 4 2.5 4 3v18c0 .5.2 1 .6 1.4c.4.4.9.6 1.4.6h12c.5 0 1-.2 1.4-.6c.4-.4.6-.9.6-1.4V3c0-.5-.2-1-.6-1.4M18 21H6V3h12v18Z"/>`),
		g.Group(children),
	)
}
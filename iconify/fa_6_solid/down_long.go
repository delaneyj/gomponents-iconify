package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownLong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M2 334.5c-3.8 8.8-2 19 4.6 26l136 144c4.5 4.8 10.8 7.5 17.4 7.5s12.9-2.7 17.4-7.5l136-144c6.6-7 8.4-17.2 4.6-26S305.5 320 296 320h-72V32c0-17.7-14.3-32-32-32h-64c-17.7 0-32 14.3-32 32v288H24c-9.6 0-18.2 5.7-22 14.5z"/>`),
		g.Group(children),
	)
}
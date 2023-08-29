package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m327.086 496.89l-142.6-142.6l-11.258-11.15l-85.6-85.6l.054-.054l11.259-11.367l85.5-85.5l.054.054l142.6-142.595L424 114.989L281.506 257.483L424 399.978ZM184.64 309.3l11.266 11.159l131.18 131.181l51.658-51.658l-142.493-142.499l142.493-142.494l-51.658-51.658l-142.392 142.394l-.054-.054l-51.813 51.812Z"/>`),
		g.Group(children),
	)
}
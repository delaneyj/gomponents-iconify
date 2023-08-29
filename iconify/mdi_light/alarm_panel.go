package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmPanel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 11v2H5v-2h3m5 0h-3v2h3v-2m5 0h-3v2h3v-2m-1-4H6v2h11V7m1-1v4H5V6h13m1-1H4c-.55 0-1 .45-1 1v13c0 .55.45 1 1 1h15c.55 0 1-.45 1-1V6c0-.55-.45-1-1-1m0-1c1.1 0 2 .9 2 2v13c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2h15M8 14H5v2h3v-2m5 0h-3v2h3v-2m5 0h-3v2h3v-2M8 17H5v2h3v-2m5 0h-3v2h3v-2m5 0h-3v2h3v-2Z"/>`),
		g.Group(children),
	)
}
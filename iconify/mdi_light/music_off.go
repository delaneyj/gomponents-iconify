package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.79 4.46l.71-.71L20.25 20.5l-.71.71l-2.42-2.43c-.35.14-.72.22-1.12.22a3 3 0 0 1-3-3c0-.4.08-.77.22-1.12L9 10.66V17a3 3 0 0 1-3 3a3 3 0 0 1-3-3a3 3 0 0 1 3-3c.77 0 1.47.29 2 .76v-5.1l-5.21-5.2M8 6.1L19 5v11c0 .53-.14 1.03-.38 1.46l-.75-.75c.08-.21.13-.46.13-.71a2 2 0 0 0-2-2c-.25 0-.5.05-.71.13l-.75-.75a3.006 3.006 0 0 1 3.46.38V9.04l-7 .78l-.92-.91L18 8.04V6.09L9 7v.84l-1-1V6.1m6 9.9a2 2 0 0 0 2 2h.31L14 15.69V16m-6 1a2 2 0 0 0-2-2a2 2 0 0 0-2 2a2 2 0 0 0 2 2a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}
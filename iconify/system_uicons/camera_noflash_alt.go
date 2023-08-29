package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraNoflashAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m7.54 5.5l1.039-1h3.92l1.92 2H16.5a2 2 0 0 1 2 2v6c0 .502-.185.961-.49 1.312m-2.51.688h-11a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h1m-2-2l14 14"/><path d="M8.306 9.454a3 3 0 0 0 4.202 4.275M13.5 11.5a3 3 0 0 0-3-3"/></g>`),
		g.Group(children),
	)
}
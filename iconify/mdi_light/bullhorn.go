package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bullhorn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 4l-3 3H3a3 3 0 0 0-3 3v3a3 3 0 0 0 3 3v3a2 2 0 0 0 2 2h1a2 2 0 0 0 2-2v-3h3l3 3h2V4h-2m.41 1H15v13h-.59l-3-3H3a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h8.41l3-3M18 7v1c1.7.25 3 1.71 3 3.47c0 1.77-1.3 3.22-3 3.47v1.01A4.504 4.504 0 0 0 18 7m0 3.06v2.83c.58-.21 1-.76 1-1.42c0-.65-.42-1.21-1-1.41M4 16h3v3a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-3Z"/>`),
		g.Group(children),
	)
}
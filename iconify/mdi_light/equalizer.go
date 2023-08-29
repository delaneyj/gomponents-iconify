package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equalizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 4v17h-3V4h3m7 6v11h-3V10h3M6 13v8H3v-8h3m8-10H9v19h5V3m7 6h-5v13h5V9M7 12H2v10h5V12Z"/>`),
		g.Group(children),
	)
}
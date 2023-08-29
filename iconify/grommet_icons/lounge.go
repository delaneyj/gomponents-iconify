package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lounge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M5 5.997C5 5.447 5.45 5 6.007 5h11.986C18.55 5 19 5.453 19 5.997V13H5V5.997ZM22 8v7.003c0 .55-.455.997-.992.997H2.992A.999.999 0 0 1 2 15.003V8m3 8v2v-2Zm14 0v2v-2Z"/>`),
		g.Group(children),
	)
}
package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleStorey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 2h14v2h-1v4h4v2h-1v10h1v2H2v-2h1V10H2V8h4V4H5V2Zm3 2v4h8V4H8Zm-3 6v10h3v-7h8v7h3V10H5Zm9 10v-5h-4v5h4ZM10.998 4.998h2.004v2.004h-2.004V4.998Z"/>`),
		g.Group(children),
	)
}
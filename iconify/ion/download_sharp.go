package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M272 160v147.37l64-64L358.63 266L256 368.63L153.37 266L176 243.37l64 64V160H92a12 12 0 0 0-12 12v296a12 12 0 0 0 12 12h328a12 12 0 0 0 12-12V172a12 12 0 0 0-12-12ZM240 32h32v128h-32z"/>`),
		g.Group(children),
	)
}
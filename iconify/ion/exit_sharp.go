package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExitSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M335.69 272h-161v-32h161V92a12 12 0 0 0-12-12h-280a12 12 0 0 0-12 12v328a12 12 0 0 0 12 12h280a12 12 0 0 0 12-12Zm83.37 0l-64 64l22.63 22.63L480.31 256L377.69 153.37L355.06 176l64 64h-83.37v32h83.37z"/>`),
		g.Group(children),
	)
}
package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateClockwiseAltFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 30H4a2.002 2.002 0 0 1-2-2V16a2.002 2.002 0 0 1 2-2h12a2.002 2.002 0 0 1 2 2v12a2.002 2.002 0 0 1-2 2zm14-15l-1.41-1.41L26 16.17V11a7.008 7.008 0 0 0-7-7h-5v2h5a5.006 5.006 0 0 1 5 5v5.17l-2.59-2.58L20 15l5 5z"/>`),
		g.Group(children),
	)
}
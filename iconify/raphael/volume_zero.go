package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.998 12.127v7.896h4.495l6.73 5.526l.003-18.95l-6.73 5.527H4.998z"/>`),
		g.Group(children),
	)
}
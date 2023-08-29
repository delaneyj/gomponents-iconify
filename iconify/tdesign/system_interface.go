package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemInterface(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 1h22v22H1V1Zm2 7.556V21h18V8.556H3Zm18-2V3H3v3.556h18ZM6 11h2.004v2.004H6V11Zm4 0h2.004v2.004H10V11Zm4 0h2.004v2.004H14V11Z"/>`),
		g.Group(children),
	)
}
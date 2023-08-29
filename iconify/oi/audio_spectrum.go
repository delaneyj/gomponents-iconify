package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioSpectrum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4 0v8h1V0H4zM2 1v6h1V1H2zm4 1v4h1V2H6zM0 3v2h1V3H0z"/>`),
		g.Group(children),
	)
}
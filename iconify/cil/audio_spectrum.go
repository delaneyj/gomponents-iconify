package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioSpectrum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 160h32v192H16zm360 0h32v192h-32zM104 88h32v328h-32zm184 8h32v320h-32zm176 0h32v320h-32zM192 16h32v480h-32z"/>`),
		g.Group(children),
	)
}
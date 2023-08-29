package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTUpOne0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="m12 29l12-12l12 12H12Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTUpOne0)"/>`),
		g.Group(children),
	)
}
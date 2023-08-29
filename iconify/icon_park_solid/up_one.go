package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUpOne0"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="m12 29l12-12l12 12H12Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUpOne0)"/>`),
		g.Group(children),
	)
}
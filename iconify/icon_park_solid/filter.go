package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFilter0"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="m6 9l14.4 16.818v12.626L27.6 42V25.818L42 9H6Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFilter0)"/>`),
		g.Group(children),
	)
}
package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Navigation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSNavigation0"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M24.5 4L9 44l15.5-9.09L40 44L24.5 4Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSNavigation0)"/>`),
		g.Group(children),
	)
}
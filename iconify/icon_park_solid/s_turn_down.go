package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func STurnDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSTurnDown0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M24 17v-4c0-4 3-7 7-7s7 3 7 7v19M10 6v29c0 4 3 7 7 7s7-3 7-7V17"/><path stroke-linecap="round" stroke-linejoin="round" d="m15 11l-5-5l-5 5"/><circle cx="38" cy="37" r="5" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSTurnDown0)"/>`),
		g.Group(children),
	)
}
package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airplane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAirplane0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20.5 10.537C20.5 6.514 22.833 4.503 24 4c1.167.503 3.5 2.514 3.5 6.537v7.543L43 31v4l-16-8v9l5 8l-8-3l-8 3l5-8v-9L5 35v-4l15.5-12.92v-7.543Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAirplane0)"/>`),
		g.Group(children),
	)
}
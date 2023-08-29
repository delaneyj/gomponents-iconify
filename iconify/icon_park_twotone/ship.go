package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ship(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTShip0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M6 20.377L24 14l18 6.377L36.667 34H11.333L6 20.377Z" clip-rule="evenodd"/><path fill="#555" d="M13 8h22l-.002 9.896L24 14l-11 3.896V8Z"/><path stroke-linecap="round" d="M24 8V4m0 20v-8M10 40l3.5 4l3.5-4l3.5 4l3.5-4l3.5 4l3.5-4l3.5 4l3.5-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTShip0)"/>`),
		g.Group(children),
	)
}
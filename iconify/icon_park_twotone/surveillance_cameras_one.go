package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurveillanceCamerasOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSurveillanceCamerasOne0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M6 26v16m0-9h7l4-6M7 20l26.474 17.649a1 1 0 0 0 1.069.025L44 32"/><path fill="#555" d="M7.078 12.719a1 1 0 0 1-.11-1.58l7.46-6.63a1 1 0 0 1 1.212-.09l27.065 17.732a1 1 0 0 1-.011 1.68l-9.144 5.82a1 1 0 0 1-1.092-.012l-25.38-16.92Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSurveillanceCamerasOne0)"/>`),
		g.Group(children),
	)
}
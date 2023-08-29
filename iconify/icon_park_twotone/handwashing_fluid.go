package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandwashingFluid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHandwashingFluid0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M24 4v7m5 6v-6H19v6M31 4H19.8C17.142 4 15 5.2 15 8m23 18.977V26a9 9 0 0 0-9-9H19a9 9 0 0 0-9 9v9a9 9 0 0 0 9 9h5"/><path fill="#555" d="M40 39.77c0 2.336-2.015 4.23-4.5 4.23S31 42.106 31 39.77c0-2.337 2.94-6.77 4.5-6.77s4.5 4.433 4.5 6.77Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHandwashingFluid0)"/>`),
		g.Group(children),
	)
}
package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TourBus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTourBus0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M9 23h30v11a2 2 0 0 1-2 2H11a2 2 0 0 1-2-2V23Z"/><path stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M9 8a2 2 0 0 1 2-2h26a2 2 0 0 1 2 2v15H9V8Z"/><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15 42a3 3 0 0 1-3-3v-3h6v3a3 3 0 0 1-3 3Zm18 0a3 3 0 0 1-3-3v-3h6v3a3 3 0 0 1-3 3Z"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M6 12v4m36-4v4"/><circle cx="15" cy="30" r="2" fill="#fff"/><circle cx="33" cy="30" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="m31 6l-9 10m16-9l-5 6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTourBus0)"/>`),
		g.Group(children),
	)
}
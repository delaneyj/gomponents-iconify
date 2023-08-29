package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ambulance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAmbulance0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linejoin="round" d="M10 35H6a2 2 0 0 1-2-2V11a2 2 0 0 1 2-2h25a2 2 0 0 1 2 2v6.899a3 3 0 0 0 1.975 2.82l7.05 2.563A3 3 0 0 1 44 26.102V33a2 2 0 0 1-2 2h-4m-20 0h12"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 19h8m-4-4v8"/><circle cx="14" cy="35" r="4" fill="#fff"/><circle cx="34" cy="35" r="4" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAmbulance0)"/>`),
		g.Group(children),
	)
}
package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParkingGarage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#61B2E4" d="M12.069 12.235h48.338v47.864H12.069z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m59.177 59.927l-46 .146a1 1 0 0 1-1.003-.997l-.147-46a1 1 0 0 1 .997-1.003l46-.146a1 1 0 0 1 1.003.997l.146 46a1 1 0 0 1-.996 1.003zM19.062 27.048L36.036 18m17.026 8.952L36.036 18"/><path d="M30.023 53.008v-22h8.389a5.48 5.48 0 0 1 5.479 5.48h0a5.48 5.48 0 0 1-5.48 5.479h-8.388"/></g>`),
		g.Group(children),
	)
}
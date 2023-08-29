package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SolarCell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92D3F5" d="M51 53H5l12-34h46z"/><path fill="#61B2E4" d="M15.25 53H51l7.92-22.44z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M51 53H5l12-34h46zM40 19L28 53m30.502-21.244L66 53"/>`),
		g.Group(children),
	)
}
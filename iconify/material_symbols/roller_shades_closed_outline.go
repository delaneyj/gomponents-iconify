package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RollerShadesClosedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21v-2h2V3h16v16h2v2h-8.25q0 .725-.513 1.238T12 22.75q-.725 0-1.238-.513T10.25 21H2Zm4-6h12V5H6v10Zm0 4h5v-2H6v2Zm7 0h5v-2h-5v2ZM6 5h12H6Z"/>`),
		g.Group(children),
	)
}
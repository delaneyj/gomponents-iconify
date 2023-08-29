package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EvStationSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 18l4-7h-2V6l-4 7.5h2Zm-4 3V3h10v9h3v7.5h2v-8.2q-.225.125-.475.162q-.25.038-.525.038q-1.05 0-1.775-.725Q15.5 10.05 15.5 9q0-.8.438-1.438q.437-.637 1.162-.912L15 4.55l1.05-1.05l3.7 3.6q.375.375.562.875q.188.5.188 1.025v12h-5v-7.5H14V21Zm14-11q.425 0 .712-.288Q19 9.425 19 9t-.288-.713Q18.425 8 18 8t-.712.287Q17 8.575 17 9t.288.712Q17.575 10 18 10Z"/>`),
		g.Group(children),
	)
}
package academicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InspireSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M48 32C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48H48zm16 64h320v320H64V96zm55.48 56.912a24.34 24.34 0 0 0-24.34 24.34a24.34 24.34 0 0 0 24.34 24.342a24.34 24.34 0 0 0 24.34-24.342a24.34 24.34 0 0 0-24.34-24.34zm54.409 6.443v200.448h44.384V227.365l80.178 132.438h46.533V159.355h-42.236v133.872L219.727 159.47l-45.838-.116zm-76.6 59.418v141.03h44.385v-141.03H97.289z"/>`),
		g.Group(children),
	)
}
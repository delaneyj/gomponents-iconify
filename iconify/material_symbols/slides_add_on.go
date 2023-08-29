package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlidesAddOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 16.95q-.825 0-1.413-.587Q3 15.775 3 14.95v-9q0-.825.587-1.413Q4.175 3.95 5 3.95h12q.825 0 1.413.587Q19 5.125 19 5.95v4.1q-.25-.05-.487-.075q-.238-.025-.513-.025t-.512.025Q17.25 10 17 10.05v-4.1H5v9h7.1q-.05.25-.075.488q-.025.237-.025.512t.025.512q.025.238.075.488Zm12 3v-3h-3v-2h3v-3h2v3h3v2h-3v3Z"/>`),
		g.Group(children),
	)
}
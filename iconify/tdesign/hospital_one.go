package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.998 1.66l10.414 9.257l-1.329 1.495L20 11.449v10.55H4V11.455l-1.094.957l-1.317-1.505L4.338 8.5l7.66-6.84ZM6 9.699V20h3v-5h6v5h3V9.67l-6-5.33l-6 5.358ZM13 20v-3h-2v3h2Zm0-13v2h2v2h-2v2h-2v-2H9V9h2V7h2Z"/>`),
		g.Group(children),
	)
}
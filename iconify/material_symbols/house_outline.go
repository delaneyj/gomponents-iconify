package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20v-8H2l10-9l4 3.6V4h3v5.3l3 2.7h-3v8h-6v-6h-2v6H5Zm2-2h2v-6h6v6h2v-7.8l-5-4.5l-5 4.5V18Zm2-6h6h-6Zm1-1.975h4q0-.8-.6-1.313T12 8.2q-.8 0-1.4.513t-.6 1.312Z"/>`),
		g.Group(children),
	)
}
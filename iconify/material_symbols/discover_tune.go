package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscoverTune(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 9V7h3V3h2v4h3v2h-8Zm3 12V11h2v10h-2ZM6 21v-4H3v-2h8v2H8v4H6Zm0-8V3h2v10H6Z"/>`),
		g.Group(children),
	)
}
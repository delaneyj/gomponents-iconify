package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExposureNegTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.1 19v-2.1l5-5.1q.825-.875 1.163-1.463T18.6 9q0-.725-.563-1.313T16.35 7.1q-.9 0-1.488.5t-.812 1.3l-2-.8q.35-1.125 1.45-2.113T16.4 5q2.075 0 3.238 1.188T20.8 9q0 1.125-.525 2.05T18.65 13.1L15 16.9l.05.1H21v2h-8.9ZM10 14H3v-2h7v2Z"/>`),
		g.Group(children),
	)
}
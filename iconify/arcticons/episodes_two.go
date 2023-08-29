package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EpisodesTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="29" x="4.5" y="11.399" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.059 11.399L7.485 7.745m5.551 3.654l3.869-6.219M24 40.4v2.42m-5.732 0h11.464M18.01 31.296a2.063 2.063 0 1 1 2.062-2.062a2.063 2.063 0 0 1-2.063 2.062Zm12.004 0a2.063 2.063 0 1 1 2.063-2.062a2.063 2.063 0 0 1-2.063 2.062Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 18.598h0A14.703 14.703 0 0 1 38.703 33.3v0a2.099 2.099 0 0 1-2.099 2.099H11.396A2.099 2.099 0 0 1 9.298 33.3v0A14.703 14.703 0 0 1 24 18.597Zm-12.697-2.199l4.177 4.917m21.217-4.917l-4.177 4.917"/>`),
		g.Group(children),
	)
}
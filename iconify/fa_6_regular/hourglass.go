package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M24 0C10.7 0 0 10.7 0 24s10.7 24 24 24h8v19c0 40.3 16 79 44.5 107.5l81.6 81.5l-81.6 81.5C48 366 32 404.7 32 445v19h-8c-13.3 0-24 10.7-24 24s10.7 24 24 24h336c13.3 0 24-10.7 24-24s-10.7-24-24-24h-8v-19c0-40.3-16-79-44.5-107.5L225.9 256l81.5-81.5C336 146 352 107.3 352 67V48h8c13.3 0 24-10.7 24-24S373.3 0 360 0H24zm168 289.9l81.5 81.5C293 391 304 417.4 304 445v19H80v-19c0-27.6 11-54 30.5-73.5l81.5-81.6zm0-67.9l-81.5-81.5C91 121 80 94.6 80 67V48h224v19c0 27.6-11 54-30.5 73.5L192 222.1z"/>`),
		g.Group(children),
	)
}
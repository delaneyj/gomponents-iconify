package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 24C0 10.7 10.7 0 24 0h336c13.3 0 24 10.7 24 24s-10.7 24-24 24h-8v19c0 40.3-16 79-44.5 107.5L225.9 256l81.5 81.5C336 366 352 404.7 352 445v19h8c13.3 0 24 10.7 24 24s-10.7 24-24 24H24c-13.3 0-24-10.7-24-24s10.7-24 24-24h8v-19c0-40.3 16-79 44.5-107.5l81.6-81.5l-81.6-81.5C48 146 32 107.3 32 67V48h-8C10.7 48 0 37.3 0 24zm110.5 347.5c-3.9 3.9-7.5 8.1-10.7 12.5h184.4c-3.2-4.4-6.8-8.6-10.7-12.5L192 289.9l-81.5 81.5zM284.2 128C297 110.4 304 89 304 67V48H80v19c0 22.1 7 43.4 19.8 61h184.4z"/>`),
		g.Group(children),
	)
}
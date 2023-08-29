package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PpeGogglesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 9C9.82 9 4 14.82 4 22a5.002 5.002 0 0 0 4 4.9V29a8 8 0 0 0 15.653 2.336a.558.558 0 0 1 .17-.264A.274.274 0 0 1 24 31c.05 0 .11.017.177.072c.07.058.135.15.17.264A8 8 0 0 0 40 29v-2.1a5.002 5.002 0 0 0 4-4.9c0-7.18-5.82-13-13-13H17Zm22.872 15.871A3.001 3.001 0 0 0 42 22c0-6.075-4.925-11-11-11H17c-6.075 0-11 4.925-11 11a3 3 0 0 0 2.128 2.871A5.002 5.002 0 0 1 13 21h22a5.002 5.002 0 0 1 4.872 3.871ZM10 26a3 3 0 0 1 3-3h22a3 3 0 0 1 3 3v3a6 6 0 0 1-11.74 1.753C25.98 29.833 25.14 29 24 29c-1.141 0-1.98.834-2.26 1.753A6 6 0 0 1 10 29v-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
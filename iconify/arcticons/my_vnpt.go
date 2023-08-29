package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyVnpt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.45 28.59c-2.36-8.807 2.174-18 10.598-21.49c8.424-3.49 18.13-.194 22.69 7.703m1.815 4.381A18.17 18.17 0 0 1 24 42.057"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.351 16.611c8.506-7.436 13.047-2.143 8.381 8.133C14.454 27.56 10.558 35 9.624 39.677c-.433 2.164 1.504 4.071 3.665 3.63C31.803 39.52 46.366 4.803 40.202 4.628c-2.89 0-6.006 1.834-7.95 3.068"/>`),
		g.Group(children),
	)
}
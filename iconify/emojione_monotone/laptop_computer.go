package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M57 53.816h-.297c.752 0 1.361-.622 1.361-1.389V19.389c0-.768-.609-1.389-1.361-1.389H7.325c-.752 0-1.362.621-1.362 1.389v33.039c0 .767.61 1.389 1.362 1.389H7l-5 4.852C2 60.334 3.348 62 4.997 62h54.007C60.559 62 62 60.338 62 58.668l-5-4.852M32.014 18.775a.777.777 0 1 1 .001 1.553a.777.777 0 0 1-.001-1.553M7.405 54.816h49.189l2.938 2.852h-21.99l-.398-.528h-10.26l-.399.528H4.466l2.939-2.852m28.911 6.022h-8.634c-.296 0-1.084 0-1.084-1.008h10.803c0 1.008-.804 1.008-1.085 1.008"/>`),
		g.Group(children),
	)
}
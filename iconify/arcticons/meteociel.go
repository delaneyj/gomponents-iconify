package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meteociel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.628 37.626c11.675.122 17.389-.015 24.84.025m0-.012a6.625 6.625 0 0 0 1.243-13.166a7.692 7.692 0 0 0-7.204-5.216a8.204 8.204 0 0 0-2.857.546A10.211 10.211 0 0 0 8.4 24.4a6.951 6.951 0 0 0 2.856 13.29m11.676-22.357a9.1 9.1 0 0 1 7.824-4.72a8.775 8.775 0 0 1 7.701 4.72a9.41 9.41 0 0 1-.186 9.315"/>`),
		g.Group(children),
	)
}
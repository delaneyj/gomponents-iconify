package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lichess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37 38.5A19.37 19.37 0 0 1 23.44 44C13.5 44 4.09 35.67 4.09 25.21c0-11 8.74-18.8 19.35-18.8a19.54 19.54 0 0 1 4.08.43A12.7 12.7 0 0 1 35 4l-1.65 5.42L43.7 27.14c1 1.42-1.37 5.84-5.71 6.48a24.59 24.59 0 0 0-4-5.07c-4.14-3.76-13.3-8.54-13.3-13.47a4 4 0 0 1 .25-1.41"/>`),
		g.Group(children),
	)
}
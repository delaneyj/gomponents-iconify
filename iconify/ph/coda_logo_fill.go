package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodaLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M136 128a40 40 0 0 0 40 40h.32c7.83.3 14-1.46 21.24-6.11A12 12 0 0 1 216 172v36a16 16 0 0 1-16 16H56a16 16 0 0 1-16-16V48a16 16 0 0 1 16-16h144a16 16 0 0 1 16 16v36a12 12 0 0 1-18.47 10.1A40.23 40.23 0 0 0 136 128Z"/>`),
		g.Group(children),
	)
}
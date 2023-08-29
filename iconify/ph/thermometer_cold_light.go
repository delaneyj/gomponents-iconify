package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerColdLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m248.29 75.82l-22.58 7.33l13.95 19.21a6 6 0 1 1-9.7 7.06L216 90.21l-14 19.21a6 6 0 1 1-9.7-7.06l13.95-19.21l-22.58-7.33a6 6 0 1 1 3.71-11.41L210 71.74V48a6 6 0 0 1 12 0v23.74l22.58-7.33a6 6 0 0 1 3.71 11.41ZM150 184a30 30 0 1 1-36-29.4V120a6 6 0 0 1 12 0v34.6a30.05 30.05 0 0 1 24 29.4Zm-12 0a18 18 0 1 0-18 18a18 18 0 0 0 18-18Zm44 0a62 62 0 1 1-100-49V48a38 38 0 0 1 76 0v87a62.06 62.06 0 0 1 24 49Zm-12 0a50.07 50.07 0 0 0-21.43-41a6 6 0 0 1-2.57-5V48a26 26 0 0 0-52 0v90a6 6 0 0 1-2.57 5A50 50 0 1 0 170 184Z"/>`),
		g.Group(children),
	)
}
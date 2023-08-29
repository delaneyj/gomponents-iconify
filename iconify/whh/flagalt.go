package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flagalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 64L768 319l256 256q0 27-19 45.5T960 639H384q-27 0-45.5-18.5T320 575V63q0-26 18.5-44.5T384 0h576q26 0 45 18.5t19 44.5v1zM384 959.5q0 26.5-18.5 45T320 1023H64q-27 0-45.5-18.5T0 959.5T18.5 914T64 895h64V63q0-26 19-44.5T192.5 0t45 18.5T256 63v832h64q27 0 45.5 19t18.5 45.5z"/>`),
		g.Group(children),
	)
}
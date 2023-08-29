package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barchart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 959.5q0 26.5-19 45.5t-45 19H64q-27 0-45.5-18.5T0 960t18.5-45.5T64 896V576q0-27 18.5-45.5T128 512h128q26 0 45 18.5t19 45.5v320h64V64q0-27 18.5-45.5T448 0h128q27 0 45.5 18.5T640 64v832h64V320q0-27 19-45.5t45-18.5h128q27 0 45.5 18.5T960 320v576q27 0 45.5 18.5t18.5 45z"/>`),
		g.Group(children),
	)
}
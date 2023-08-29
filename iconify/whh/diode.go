package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 385H832v192q0 26-19 45t-45 19t-45-19t-19-45V385h-10L296 634q-17 17-40-8V385H64q-26 0-45-19T0 320.5t19-45T64 257h192V15q24-25 40-8l399 250h9V65q0-27 18.5-45.5T768 1t45.5 18.5T832 65v192h128q26 0 45 18.5t19 45t-19 45.5t-45 19zM320 129v384l320-192z"/>`),
		g.Group(children),
	)
}
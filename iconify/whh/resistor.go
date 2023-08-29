package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Resistor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 376H818l-47 168q-3 12-13 18.5t-22 5.5q-12 1-22-5.5T701 544l-93-335l-93 335q-3 12-13 18.5t-22 5.5q-12 1-22-5.5T445 544l-93-335l-32 167H64q-26 0-45-18.5t-19-45T19 267t45-19h206l47-168q3-11 13-17.5t22-5.5q12-1 22 5.5T387 80l93 336l93-336q3-11 13-17.5t22-5.5q12-1 22 5.5T643 80l93 336l32-168h192q27 0 45.5 19t18.5 45.5t-18.5 45T960 376z"/>`),
		g.Group(children),
	)
}
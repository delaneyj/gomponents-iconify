package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pigpenz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M977 1008.5q-34 22.5-72.5 11T846 969L513 324L180 969q-20 39-58.5 50.5t-72.5-11T4.5 943t9.5-81L430 55q0-1 1-2.5l1-1.5q13-23 33-36q34-22 72.5-10.5T596 55l416 807q20 38 9.5 81t-44.5 65.5zM513 768q53 0 90.5 37.5T641 896t-37.5 90.5T513 1024t-90.5-37.5T385 896t37.5-90.5T513 768z"/>`),
		g.Group(children),
	)
}
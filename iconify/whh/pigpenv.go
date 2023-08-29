package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pigpenv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M977 1008.5q-34 22.5-72.5 11T846 969L513 323L180 969q-20 39-58.5 50.5t-72.5-11T4.5 943t9.5-81L430 54l1-1.5l1-1.5q13-23 33-36q34-22 72.5-10.5T596 54l416 808q20 38 9.5 81t-44.5 65.5z"/>`),
		g.Group(children),
	)
}
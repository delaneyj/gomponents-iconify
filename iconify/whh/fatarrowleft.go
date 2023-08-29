package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fatarrowleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 1025H383v-1q-41 0-70-39L29 607Q0 567 0 512t29-94L313 39q29-38 70-39h513q53 0 90.5 37.5T1024 128v769q0 53-37.5 90.5T896 1025z"/>`),
		g.Group(children),
	)
}
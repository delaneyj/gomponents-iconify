package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Logstash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.802 9.599V32C9.864 32 2.401 26.667 2.401 19.599V0h4.797c5.068 0 9.604 4.531 9.604 9.599zm3.198 8V32h9.599V17.599z"/>`),
		g.Group(children),
	)
}
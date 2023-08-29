package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bullseye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm0 14.9c-3.8 0-6.9-3.1-6.9-6.9S4.2 1.1 8 1.1s6.9 3.1 6.9 6.9s-3.1 6.9-6.9 6.9z"/><path fill="currentColor" d="M8 2.3C4.8 2.3 2.3 4.8 2.3 8s2.6 5.7 5.7 5.7s5.7-2.6 5.7-5.7S11.2 2.3 8 2.3zm0 10.3c-2.5 0-4.6-2.1-4.6-4.6S5.5 3.4 8 3.4s4.6 2.1 4.6 4.6c0 2.5-2.1 4.6-4.6 4.6z"/><path fill="currentColor" d="M8 4.6C6.1 4.6 4.6 6.1 4.6 8s1.5 3.4 3.4 3.4s3.4-1.5 3.4-3.4S9.9 4.6 8 4.6z"/>`),
		g.Group(children),
	)
}
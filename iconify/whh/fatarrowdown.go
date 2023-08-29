package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fatarrowdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1023.996 640q0 41-39 70l-378 284q-39 29-94.5 29t-94.5-29l-378-284q-39-29-39-70h-1V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v512z"/>`),
		g.Group(children),
	)
}
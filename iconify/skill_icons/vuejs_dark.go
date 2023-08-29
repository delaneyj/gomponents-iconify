package skill_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VuejsDark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="none"><rect width="256" height="256" fill="#242938" rx="60"/><path fill="#41B883" d="M182 50h36l-90 155.25L38 50h68.85L128 86l20.7-36H182Z"/><path fill="#41B883" d="m38 50l90 155.25L218 50h-36l-54 93.15L73.55 50H38Z"/><path fill="#fff" d="M73.55 50L128 143.6L182 50h-33.3L128 86l-21.15-36h-33.3Z"/></g>`),
		g.Group(children),
	)
}
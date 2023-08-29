package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unavailable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18.364 5.636A9 9 0 1 1 5.636 18.364A9 9 0 0 1 18.364 5.636Zm-2.172 11.97L6.393 7.808a7.001 7.001 0 0 0 9.8 9.8ZM16.95 7.05a7.002 7.002 0 0 1 .657 9.142l-9.8-9.799a7.001 7.001 0 0 1 9.143.657Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
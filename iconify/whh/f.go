package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func F(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M705 128H160q-13 0-22.5 9.5T128 160v256q0 13 9.5 22.5T160 448h416q27 0 46 19t19 45.5t-19 45t-46 18.5H160q-13 0-22.5 9.5T128 608v352q0 27-18.5 45.5T64 1024t-45.5-18.5T0 960V64q0-26 18.5-45T64 0h641q26 0 45 19t19 45.5t-19 45t-45 18.5z"/>`),
		g.Group(children),
	)
}
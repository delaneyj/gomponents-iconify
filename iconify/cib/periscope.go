package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Periscope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.183 32c2.5 0 12.416-11.251 12.416-18.797C28.599 5.984 22.959 0 16.183 0C9.042 0 3.402 5.984 3.402 13.203C3.402 20.746 13.318 32 16.183 32zM14.188 5.131a8.363 8.363 0 0 1 2.031-.251c4 0 7.437 3.401 7.437 7.823c0 3.98-3.437 7.38-7.459 7.38c-4.457 0-7.9-3.4-7.9-7.38c0-1.864.599-3.52 1.64-4.843v.041c0 1.661 1.339 2.979 3 2.979S16 9.521 16 7.86a2.98 2.98 0 0 0-1.812-2.735z"/>`),
		g.Group(children),
	)
}
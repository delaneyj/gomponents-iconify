package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spectrum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M0 13.599A2.398 2.398 0 0 0 2.401 16h2.401c6.181 0 11.197 5.016 11.197 11.197v2.401a2.398 2.398 0 0 0 2.401 2.401h11.197a2.398 2.398 0 0 0 2.401-2.401v-2.401C31.998 12.177 19.821 0 4.801 0H2.4A2.398 2.398 0 0 0-.001 2.401z"/>`),
		g.Group(children),
	)
}
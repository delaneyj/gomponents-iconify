package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21v-3c-3.36-.27-6-3.08-6-6.5V11h1v.5a5.5 5.5 0 0 0 5.5 5.5a5.5 5.5 0 0 0 5.5-5.5V11h1v.5c0 3.42-2.64 6.23-6 6.5v3h-1m.5-18A2.5 2.5 0 0 1 14 5.5v6a2.5 2.5 0 0 1-2.5 2.5A2.5 2.5 0 0 1 9 11.5v-6A2.5 2.5 0 0 1 11.5 3m0 1A1.5 1.5 0 0 0 10 5.5v6a1.5 1.5 0 0 0 1.5 1.5a1.5 1.5 0 0 0 1.5-1.5v-6A1.5 1.5 0 0 0 11.5 4Z"/>`),
		g.Group(children),
	)
}
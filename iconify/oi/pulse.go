package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pulse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="m3.25 0l-.47 1.53L2 4.09l-.03-.06l-.09-.34H0v1h1.16l.38 1.16l.47 1.47l.47-1.5l.78-2.5l.78 2.5l.41 1.34l.53-1.28l.59-1.47l.13.28h2.31v-1H6.32l-.38-.75l-.5-.97L5.03 3l-.47 1.19l-.84-2.66L3.25 0z"/>`),
		g.Group(children),
	)
}
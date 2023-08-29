package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbOnTwenty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 11h3v2H1v-2m3.9-7.5L3.5 4.9L5.6 7L7 5.6L4.9 3.5M13 1h-2v3h2V1m7 10v2h3v-2h-3M10 22c0 .6.4 1 1 1h2c.6 0 1-.4 1-1v-1h-4v1m9.1-18.5L17 5.6L18.4 7l2.1-2.1l-1.4-1.4M18 12c0 2.2-1.2 4.2-3 5.2V19c0 .6-.4 1-1 1h-4c-.6 0-1-.4-1-1v-1.8c-1.8-1-3-3-3-5.2c0-3.3 2.7-6 6-6s6 2.7 6 6m-2 0c0-2.21-1.79-4-4-4s-4 1.79-4 4s1.79 4 4 4s4-1.79 4-4Z"/>`),
		g.Group(children),
	)
}
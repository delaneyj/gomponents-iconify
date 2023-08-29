package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbOnSeventy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 5.6L5.6 7L3.5 4.9l1.4-1.4L7 5.6M13 1h-2v3h2V1M1 13h3v-2H1v2m9 9c0 .6.4 1 1 1h2c.6 0 1-.4 1-1v-1h-4v1m9.1-18.5L17 5.6L18.4 7l2.1-2.1l-1.4-1.4M18 12c0 2.2-1.2 4.2-3 5.2V19c0 .6-.4 1-1 1h-4c-.6 0-1-.4-1-1v-1.8c-1.8-1-3-3-3-5.2c0-3.3 2.7-6 6-6s6 2.7 6 6m-9.86-1h7.72c-.45-1.72-2-3-3.86-3s-3.41 1.28-3.86 3M20 11v2h3v-2h-3Z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbOnEighty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 5.6L5.6 7L3.5 4.9l1.4-1.4L7 5.6M1 13h3v-2H1v2M13 1h-2v3h2V1m-3 21c0 .6.4 1 1 1h2c.6 0 1-.4 1-1v-1h-4v1m10-11v2h3v-2h-3m-.9-7.5L17 5.6L18.4 7l2.1-2.1l-1.4-1.4M18 12c0 2.2-1.2 4.2-3 5.2V19c0 .6-.4 1-1 1h-4c-.6 0-1-.4-1-1v-1.8c-1.8-1-3-3-3-5.2c0-3.3 2.7-6 6-6s6 2.7 6 6m-9.44-2h6.88C14.75 8.81 13.5 8 12 8s-2.75.81-3.44 2Z"/>`),
		g.Group(children),
	)
}
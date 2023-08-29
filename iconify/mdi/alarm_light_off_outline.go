package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmLightOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 7.3L9 5.8c.9-.5 1.9-.8 3-.8c3.3 0 6 2.7 6 6v3.8l-2-2V11c0-2.2-1.8-4-4-4c-.5 0-1 .1-1.5.3m9.6-2.5l-1.4-1.4l-2.1 2.1L18 6.9l2.1-2.1m-.6 5.7v2h3v-2h-3m-15 0h-3v2h3v-2M2.4 1.7l19.7 19.7l-1.3 1.3l-.7-.7H4c0-1.1.9-2 2-2h12.1l-1-1H6v-8c0-.9.2-1.8.6-2.5L1.1 3l1.3-1.3M8 17h7.1l-7-7c0 .3-.1.7-.1 1v6m5-16h-2v3h2V1Z"/>`),
		g.Group(children),
	)
}
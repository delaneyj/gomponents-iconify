package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M10.9 7.875c.664.266 1.1.91 1.1 1.625v13a1.75 1.75 0 1 1-3.5 0v-8.625l-.157.165a1.75 1.75 0 1 1-2.535-2.413l3.175-3.334a1.75 1.75 0 0 1 1.917-.418ZM20 7.75a6 6 0 0 0-6 6v4.5a6 6 0 0 0 12 0v-4.5a6 6 0 0 0-6-6Zm-2.5 6a2.5 2.5 0 0 1 5 0v4.5a2.5 2.5 0 0 1-5 0v-4.5Z"/><path d="M1 6a5 5 0 0 1 5-5h20a5 5 0 0 1 5 5v20a5 5 0 0 1-5 5H6a5 5 0 0 1-5-5V6Zm5-3a3 3 0 0 0-3 3v20a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6Z"/></g>`),
		g.Group(children),
	)
}
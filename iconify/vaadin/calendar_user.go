package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 0h1v3H3V0zm8 0h1v3h-1V0z"/><path fill="currentColor" d="M9 14.1c0-.1 0-.1 0 0L1 14V6h13v1.2c.4.1.7.3 1 .6V1h-2v3h-3V1H5v3H2V1H0v14h9v-.9z"/><path fill="currentColor" d="M15 10a2 2 0 1 1-3.999.001A2 2 0 0 1 15 10z"/><path fill="currentColor" d="M13.9 12h-1.8c-1.1 0-2.1.9-2.1 2.1V16h6v-1.9c0-1.2-.9-2.1-2.1-2.1z"/>`),
		g.Group(children),
	)
}
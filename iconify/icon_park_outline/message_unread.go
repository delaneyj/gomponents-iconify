package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageUnread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 16v20H29l-5 5l-5-5H4V6h30"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M23 20h2.003m7.998 0H35m-21.999 0H15"/><circle cx="43" cy="7" r="3" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
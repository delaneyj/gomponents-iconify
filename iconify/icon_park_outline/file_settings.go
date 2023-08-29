package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2ZM30 4l10 10"/><circle cx="24" cy="27" r="5"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 19v3m0 10v3m5.828-14l-2.121 2.121M19.828 31l-2.121 2.121M18 21l2.121 2.121M28 31l2.121 2.121M16 27h3m10 0h3"/></g>`),
		g.Group(children),
	)
}
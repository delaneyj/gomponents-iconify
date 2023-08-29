package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Move(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#9B9B9A" d="M11 36.323h24.698v24.698H11z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M56 38.219v2.5h-2.5"/><path stroke-dasharray="3.94 3.94" d="M49.56 40.719H35.772m-4.47-6.44V20.491"/><path d="M31.302 18.521v-2.5h2.5"/><path stroke-dasharray="3.94 3.94" d="M37.742 16.021H51.53"/><path d="M53.5 16.021H56v2.5"/><path stroke-dasharray="3.94 3.94" d="M56 22.46v13.789"/><path stroke-miterlimit="10" d="M11 36.323h24.698v24.698H11z"/><path stroke-miterlimit="10" d="m35.601 36.419l7.033-7.032V36m0-6.613h-6.571"/></g>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.9 14L14 41.9c-.8.8-2 .8-2.8 0l-5.1-5.1c-.8-.8-.8-2 0-2.8L34 6.1c.8-.8 2-.8 2.8 0l5.1 5.1c.8.8.8 2.1 0 2.8ZM20 20l3.7 3.7M7 33.1l3.6 3.7m.7-8.1l2.1 2.1m2.3-6.4l2 2M33.1 7l3.7 3.6m-12.4 5.1l2 2m2.3-6.4l2.1 2.1"/>`),
		g.Group(children),
	)
}
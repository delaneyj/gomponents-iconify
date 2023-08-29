package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningThreeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="2" transform="rotate(180 12 12)"/><circle cx="20" cy="14" r="2" transform="rotate(180 20 14)"/><circle cx="2" cy="2" r="2" transform="matrix(-1 0 0 1 6 8)"/><path stroke-linecap="round" d="M4 12v2.75M4 19v-1.5M20 12V9.25M20 5v1.5M12 19v-5m8 5v-3M4 5v2.667M12 10V7.5M12 5v.5"/></g>`),
		g.Group(children),
	)
}
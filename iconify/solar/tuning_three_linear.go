package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningThreeLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="2" transform="rotate(180 12 12)"/><circle cx="20" cy="14" r="2" transform="rotate(180 20 14)"/><circle cx="2" cy="2" r="2" transform="matrix(-1 0 0 1 6 8)"/><path stroke-linecap="round" d="M12 8V5m8 5V5M4 14v5m8 0v-3m8 3v-1M4 5v1"/></g>`),
		g.Group(children),
	)
}
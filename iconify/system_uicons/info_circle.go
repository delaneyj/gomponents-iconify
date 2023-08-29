package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" transform="translate(2 2)"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="8.5" cy="8.5" r="8"/><path d="M8.5 12.5v-4h-1m0 4h2"/></g><circle cx="8.5" cy="5.5" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Steoller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 24H12l2.5 10H36l4-10Zm-28 0l-4-9H3.5"/><circle cx="20" cy="41" r="3"/><circle cx="31" cy="41" r="3"/><path d="m23 8l5 16h12s3.5-8-2.5-14S25.667 6.667 23 8Z"/></g>`),
		g.Group(children),
	)
}
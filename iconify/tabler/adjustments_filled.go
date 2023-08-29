package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustmentsFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M6 3a1 1 0 0 1 .993.883L7 4v3.171a3.001 3.001 0 0 1 0 5.658V20a1 1 0 0 1-1.993.117L5 20v-7.17a3.002 3.002 0 0 1-1.995-2.654L3 10l.005-.176A3.002 3.002 0 0 1 5 7.17V4a1 1 0 0 1 1-1zm6 0a1 1 0 0 1 .993.883L13 4v9.171a3.001 3.001 0 0 1 0 5.658V20a1 1 0 0 1-1.993.117L11 20v-1.17a3.002 3.002 0 0 1-1.995-2.654L9 16l.005-.176A3.002 3.002 0 0 1 11 13.17V4a1 1 0 0 1 1-1zm6 0a1 1 0 0 1 .993.883L19 4v.171a3.001 3.001 0 0 1 0 5.658V20a1 1 0 0 1-1.993.117L17 20V9.83a3.002 3.002 0 0 1-1.995-2.654L15 7l.005-.176A3.002 3.002 0 0 1 17 4.17V4a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}
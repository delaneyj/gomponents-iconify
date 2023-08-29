package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M11.99 1.968c1.023 0 1.97.521 2.512 1.359l.103.172l7.1 12.25l.062.126a3 3 0 0 1-2.568 4.117L19 20H5l-.049-.003l-.112.002a3 3 0 0 1-2.268-1.226l-.109-.16a3 3 0 0 1-.32-2.545l.072-.194l.06-.125L9.366 3.516a3 3 0 0 1 2.625-1.548z"/></g>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cuvette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M18 10h12v17l14 14v3H4v-3l14-14V10Z" clip-rule="evenodd"/><path d="M40.5 37.5L37 34l-3.5-3.5m7 7L37 34l-3.5-3.5m-19 0L11 34l-3.5 3.5m7-7L11 34l-3.5 3.5M18 4h12m-6 23v1m0-10v3M10 35h28"/></g>`),
		g.Group(children),
	)
}
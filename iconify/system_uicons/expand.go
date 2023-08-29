package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 7.5v-5h-5m5 0l-6 5.929M7.5 18.5l-5 .023V13.5m6-1l-6 6"/>`),
		g.Group(children),
	)
}
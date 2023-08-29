package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EndTimeSort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 5V30.0036H42V5"/><path stroke-linejoin="round" d="M30 37L24 43L18 37"/><path stroke-linejoin="round" d="M24 30V43"/><path stroke-linejoin="round" d="M18.9604 10.9786L23.9972 15.9928L18.9604 21.0903"/><path d="M29 10.002V22.0001"/></g>`),
		g.Group(children),
	)
}
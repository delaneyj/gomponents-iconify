package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotationLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M13 29L16.6667 21M35 29L31.3333 21M31.3333 21L29.5 17L24 5L18.5 17L16.6667 21M31.3333 21H16.6667"/><path d="M42 37H6L12 43"/></g>`),
		g.Group(children),
	)
}
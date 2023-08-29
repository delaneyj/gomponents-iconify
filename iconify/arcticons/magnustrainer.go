package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magnustrainer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.17 22.84h15.66m-2.58 0h-10.5L14.62 43.5h18.76l-4.13-20.66zM12.55 43.5h22.9m-22.9-3.62h22.9m-6.2-17.04l4.64-9.27a22 22 0 0 0-19.78 0l4.64 9.27ZM24 11.22V4.5m-3.36 2.93h6.72"/>`),
		g.Group(children),
	)
}
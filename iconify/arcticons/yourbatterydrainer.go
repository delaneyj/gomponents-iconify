package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yourbatterydrainer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.694 6.918L42.04 27.264a1.278 1.278 0 0 1-.011 1.976l-9.982 11.944a1.3 1.3 0 0 1-1.797-.121L5.954 20.759a1.243 1.243 0 0 1-.158-1.773L15.82 6.992a1.198 1.198 0 0 1 1.874-.074Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.375 10.854l3.406 2.846l-8.283 9.91l-3.405-2.845Zm5.536 4.583l3.405 2.846l-8.282 9.911l-3.405-2.846Zm5.583 4.58l3.406 2.845l-8.283 9.911l-3.405-2.846Zm10.164 12.161l2.642 2.21l-5.139 6.15l-2.643-2.21Z"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Serviigo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.47 22.95a10.165 10.165 0 0 0-15.855 8.422m4.481 8.428a10.165 10.165 0 0 0 15.847-8.428"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.185 31.372h3.82l-5.253 5.252L4.5 31.372h6.685zm17.189-.001h-3.819l5.252-5.253l5.252 5.253h-6.685zm5.28-5.125a14.809 14.809 0 0 0-11.328-9.443m16.039 7.703a19.837 19.837 0 0 0-15.175-12.65M43.5 22.609A25.318 25.318 0 0 0 24.132 6.464"/>`),
		g.Group(children),
	)
}
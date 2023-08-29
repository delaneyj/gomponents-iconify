package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfinityAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.012 27.797a3.68 3.68 0 0 0 4.196 0"/><path fill="currentColor" d="M30.245 23.366a.752.752 0 0 1 .766.75a.758.758 0 0 1-1.516 0a.748.748 0 0 1 .75-.75Zm-7.655-.116a.752.752 0 0 1 .767.75a.758.758 0 0 1-1.516 0a.748.748 0 0 1 .75-.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.503 7.41a4.103 4.103 0 1 1-4.107 4.107a4.102 4.102 0 0 1 4.107-4.107Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.467 9.313c4.57-.045 2.309 2.863-3.868 4.54s-10.516.427-6.614-1.83m32.931 16.724c-2.253-13.92-22.87-16.777-25.976-.417m-.39 1.579a42.762 42.762 0 0 0 26.668.248c6.463 1.688 5.5 4.212-2.15 5.638s-19.09 1.214-25.554-.474C5.43 33.732 5.88 31.38 12.55 29.91Zm17.397 8.232a4.063 4.063 0 0 1-7.452.01m-4.551-.418a3.69 3.69 0 0 1-6.618-1.392m29.427-.116a3.478 3.478 0 0 1-6.22 1.41"/>`),
		g.Group(children),
	)
}
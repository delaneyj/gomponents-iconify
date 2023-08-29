package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uchealth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 21.637v2.73a1.66 1.66 0 0 0 1.655 1.655h0a1.66 1.66 0 0 0 1.655-1.655v-2.73m0 2.73v1.655m6.902-6.619v6.619m0-2.73a1.66 1.66 0 0 1 1.654-1.655h0a1.66 1.66 0 0 1 1.655 1.655v2.73m18.169-6.619v6.619m0-2.73a1.66 1.66 0 0 1 1.655-1.655h0a1.66 1.66 0 0 1 1.655 1.655v2.73m-6.041-5.792v5.792m-.911-4.385h1.738m-11.269 3.558a1.601 1.601 0 0 1-1.406.827h0a1.66 1.66 0 0 1-1.655-1.655v-1.075a1.66 1.66 0 0 1 1.655-1.655h0a1.66 1.66 0 0 1 1.655 1.655v.579h-3.31m10.501-4.468v5.792a.782.782 0 0 0 .828.827h.248m-18.785-.827a1.602 1.602 0 0 1-1.407.827h0a1.66 1.66 0 0 1-1.655-1.655v-1.075a1.66 1.66 0 0 1 1.655-1.655h0a1.602 1.602 0 0 1 1.407.827m15.748 1.903a1.66 1.66 0 0 1-1.655 1.655h0a1.66 1.66 0 0 1-1.655-1.655v-1.075a1.66 1.66 0 0 1 1.655-1.655h0a1.66 1.66 0 0 1 1.655 1.655m0 2.73v-4.385M4.5 29h39"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}
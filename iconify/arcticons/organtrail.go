package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Organtrail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.5 4.5h5.57v1.86h1.86v3.71h1.86v1.86h-7.43v1.86h3.71v1.85h11.14v1.86h1.86v3.71h-3.71v-1.85h-9.29v5.57h-3.71v1.86h1.85v1.85h1.86v5.57h1.86v5.58h3.71v1.85h-3.71v1.86h-3.72v-9.29H17.5v3.72h-1.86v3.71h3.72v1.86h-5.57v-1.86h-1.86v-3.71h1.86v-3.72h1.85V17.5h1.86v-3.71h-1.86V6.36h1.86Zm3.71 3.71v1.86h1.86V8.21Z"/>`),
		g.Group(children),
	)
}
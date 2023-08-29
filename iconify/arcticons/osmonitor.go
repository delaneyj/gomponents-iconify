package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Osmonitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.74h37a1 1 0 0 1 1 1v31.72a1 1 0 0 1-1 1h-37a1 1 0 0 1-1-1V6.74a1 1 0 0 1 1-1Z"/><circle cx="40.27" cy="37.16" r=".75" fill="currentColor"/><circle cx="35.79" cy="37.16" r=".75" fill="currentColor"/><circle cx="31.31" cy="37.16" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.32 8.17h33.36v26.06H7.32z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.32 22.6h6.49l2.53-8.61l4.29 15.79L24 17.97l2.18 7.68L28 22.6h12.68"/>`),
		g.Group(children),
	)
}
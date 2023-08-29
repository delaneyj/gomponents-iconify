package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 7.15L4.5 22.4l2.588 3.31l2.997-2.345v17.486h27.83V23.365l2.997 2.344L43.5 22.4l-5.585-4.367v-6.145h-3.103v3.652L24 7.149Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.009 40.85V29.064A5.009 5.009 0 0 0 24 24.055h0a5.009 5.009 0 0 0-5.009 5.008v11.788"/>`),
		g.Group(children),
	)
}
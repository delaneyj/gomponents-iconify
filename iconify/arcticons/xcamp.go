package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xcamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5c-7.66.001-13.868 6.21-13.87 13.87c.155 3.587 1.105 5.784 2.894 8.426L24.044 43.5l11.053-16.868c1.655-2.341 2.699-5.564 2.772-8.263C37.87 10.71 31.66 4.501 24 4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.7 19.22v-1.213h-3.635v2.424a3.634 3.634 0 0 0 3.634 3.634h12.115v-3.634H20.91a1.211 1.211 0 0 1-1.21-1.212Zm8.48-2.165v.352h3.635v-1.564a3.634 3.634 0 0 0-3.635-3.634H16.065v3.634h10.904c.669 0 1.21.543 1.21 1.212Z"/>`),
		g.Group(children),
	)
}
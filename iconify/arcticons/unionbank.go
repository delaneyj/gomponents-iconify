package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unionbank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 12.993v15.3c0 4.185 3.341 7.148 7.133 6.661c3.47-.446 6.074-3.67 6.074-7.234V14.69m0 .001l6.41.824c1.91.246 3.383 2.334 3.383 4.676s-1.473 4.241-3.383 4.241h0c1.91 0 3.383 1.899 3.383 4.241s-1.473 4.43-3.383 4.676l-4.8.617m4.8-9.534h-6.41"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AliorMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 31.494s7.952.163 11.804-1.782c1.363-.689 1.573-1.364 1.573-1.364c.18-2.634 1.283-10.79-2.553-14.903C31.141 9.495 27.117 9.72 24 9.54c-3.117.18-7.14-.046-10.824 3.904c-3.836 4.113-2.733 12.27-2.552 14.903c0 0 .21.675 1.572 1.364C16.048 31.657 24 31.494 24 31.494Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.558 20.817C36.989 26 28.056 26.087 24 26.087S11.01 26 10.442 20.816"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.658 23.657c1.77-.143 4.055-.698 5.13.566c.843.992 1.818 4.693-2.76 8.335c-4.36 3.47-11.015 4.062-16.028 4.153c-5.013-.09-11.667-.683-16.029-4.153c-4.577-3.642-3.602-7.343-2.76-8.335c1.076-1.264 3.36-.709 5.13-.566"/>`),
		g.Group(children),
	)
}
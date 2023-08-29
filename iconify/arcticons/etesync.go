package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Etesync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 3.5l-6 6H9.51V18l-6 6l6 6v8.48H18l6 6l6-6h8.48V30l6-6l-6-6V9.51H30Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.67 29.18A9.26 9.26 0 0 0 24 14.75m-7.67 4.08A9.16 9.16 0 0 0 14.75 24A9.25 9.25 0 0 0 24 33.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 16.18v3.48l-4.78-4.79L24 10.09h0v6.09zm0 15.64v-3.47l4.78 4.78L24 37.91v-6.09z"/>`),
		g.Group(children),
	)
}
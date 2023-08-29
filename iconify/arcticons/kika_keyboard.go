package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KikaKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.111 31.912L24 43.5L8.889 31.912a3.138 3.138 0 0 1-1.229-2.49V7.01a2.51 2.51 0 0 1 2.511-2.51h27.658a2.51 2.51 0 0 1 2.51 2.51h0v22.412c0 .976-.453 1.896-1.228 2.49Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 7.751c-7.181 0-13 5.819-13 13s5.819 13 13 13s13-5.818 13-13s-5.819-13-13-13Zm5.965 12.073l-4.193-2.517a.633.633 0 0 1 0-1.08l4.193-2.515m-11.93 0l4.193 2.516a.629.629 0 0 1 0 1.08l-4.193 2.516m12.916 4.35a6.64 6.64 0 0 1-6.583 5.803h-.736a6.64 6.64 0 0 1-6.583-5.804a.74.74 0 0 1 .732-.83H30.22c.442 0 .787.388.732.83Z"/>`),
		g.Group(children),
	)
}
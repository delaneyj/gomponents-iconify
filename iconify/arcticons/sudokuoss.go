package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sudokuoss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.41 5.5v37m12.09-37v37M10.033 14.627a2.78 2.78 0 0 0 2.1.61h.27a2.32 2.32 0 0 0 2.28-2.28h0a2.32 2.32 0 0 0-2.28-2.28h-2.37v-2.46h4.65M42.5 18h-37m37 12.46h-37m19.561-2.362a2.006 2.006 0 0 0 2-2h0a2.006 2.006 0 0 0-2-2h0a2.006 2.006 0 0 0 2-2h0a2.006 2.006 0 0 0-2-2m-3.3 7.4a3.253 3.253 0 0 0 2.5.7h.8m-3.3-7.4a3.402 3.402 0 0 1 2.5-.7h.8"/><circle cx="37.068" cy="35.611" r="2.7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.668 40.011a2.456 2.456 0 0 0 2.2 1h.2a2.689 2.689 0 0 0 2.7-2.7v-2.7M21.924 10.39a2.67 2.67 0 0 1 2.6-2.7a2.687 2.687 0 0 1 1.9 4.6c-1.1.9-4.5 3.5-4.5 3.5h5.3"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Merck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.15 14.995c-1.783 0-7.204 8.858-14.15 8.858s-12.367-8.858-14.15-8.858H4.5v15.006a3.012 3.012 0 0 0 3.004 3.004H9.85v-5.163c0-2.018.986-3.05 2.992-3.05c3.203 0 7.474 7.227 11.158 7.227s7.955-7.227 11.158-7.227c2.006 0 2.992 1.032 2.992 3.05v2.16a3.004 3.004 0 0 0 3.003 3.003H43.5V17.103a2.108 2.108 0 0 0-2.108-2.108H38.15Z"/>`),
		g.Group(children),
	)
}
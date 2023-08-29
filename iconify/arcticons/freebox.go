package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freebox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.9 22.52V11.31h-3.82m8.02 0H37v11.21H26.1M37 16.91h-6.54M11 25.48h10.9v11.21H11m15.1-11.21H37v11.21m-15.1-5.6h-6.54"/>`),
		g.Group(children),
	)
}